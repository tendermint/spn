package integration_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/starport/starport/pkg/cmdrunner"
	"github.com/tendermint/starport/starport/pkg/cmdrunner/step"
	"github.com/tendermint/starport/starport/pkg/httpstatuschecker"
	"github.com/tendermint/starport/starport/pkg/xurl"
)

const (
	serveTimeout = time.Minute * 5
)

var isCI, _ = strconv.ParseBool(os.Getenv("CI"))

// env provides an isolated testing environment and what's needed to
// make it possible.
type env struct {
	t   *testing.T
	ctx context.Context
}

// env creates a new testing environment.
func newEnv(t *testing.T) env {
	ctx, cancel := context.WithCancel(context.Background())
	e := env{
		t:   t,
		ctx: ctx,
	}
	t.Cleanup(cancel)
	return e
}

// Ctx returns parent context for the test suite to use for cancelations.
func (e env) Ctx() context.Context {
	return e.ctx
}

type execOptions struct {
	ctx                    context.Context
	shouldErr, shouldRetry bool
	stdout, stderr         io.Writer
}

type execOption func(*execOptions)

// ExecShouldError sets the expectations of a command's execution to end with a failure.
func ExecShouldError() execOption {
	return func(o *execOptions) {
		o.shouldErr = true
	}
}

// ExecCtx sets cancelation context for the execution.
func ExecCtx(ctx context.Context) execOption {
	return func(o *execOptions) {
		o.ctx = ctx
	}
}

// ExecStdout captures stdout of an execution.
func ExecStdout(w io.Writer) execOption {
	return func(o *execOptions) {
		o.stdout = w
	}
}

// ExecSterr captures stderr of an execution.
func ExecStderr(w io.Writer) execOption {
	return func(o *execOptions) {
		o.stderr = w
	}
}

// ExecRetry retries command until it is successful before context is canceled.
func ExecRetry() execOption {
	return func(o *execOptions) {
		o.shouldRetry = true
	}
}

// Exec executes a command step with options where msg describes the expectation from the test.
// unless calling with Must(), Exec() will not exit test runtime on failure.
func (e env) Exec(msg string, step *step.Step, options ...execOption) (ok bool) {
	opts := &execOptions{
		ctx:    e.ctx,
		stdout: ioutil.Discard,
		stderr: ioutil.Discard,
	}
	for _, o := range options {
		o(opts)
	}
	var (
		stdout = &bytes.Buffer{}
		stderr = &bytes.Buffer{}
	)
	copts := []cmdrunner.Option{
		cmdrunner.DefaultStdout(io.MultiWriter(stdout, opts.stdout)),
		cmdrunner.DefaultStderr(io.MultiWriter(stderr, opts.stderr)),
	}
	if isCI {
		copts = append(copts, cmdrunner.EndSignal(os.Kill))
	}
	err := cmdrunner.
		New(copts...).
		Run(opts.ctx, step)
	if err == context.Canceled {
		err = nil
	}
	if err != nil && opts.shouldRetry && opts.ctx.Err() == nil {
		time.Sleep(time.Second)
		return e.Exec(msg, step, options...)
	}
	if err != nil {
		msg = fmt.Sprintf("%s\n\nLogs:\n\n%s\n\nError Logs:\n\n%s\n",
			msg,
			stdout.String(),
			stderr.String())
	}
	if opts.shouldErr {
		return assert.Error(e.t, err, msg)
	}
	return assert.NoError(e.t, err, msg)
}

// Serve serves an application lives under path with options where msg describes the
// expection from the serving action.
// unless calling with Must(), Serve() will not exit test runtime on failure.
func (e env) Serve(msg string, path string, options ...execOption) (ok bool) {
	return e.Exec(msg,
		step.New(
			step.Exec(
				"starport",
				"serve",
				"-v",
			),
			step.Stderr(os.Stderr),
			step.Stdout(os.Stdout),
			step.Workdir(path),
		),
		options...,
	)
}

// IsAppServed checks that app is served properly and servers are started to listening
// before ctx canceled.
func (e env) IsAppServed(ctx context.Context) error {
	checkAlive := func() error {
		ok, err := httpstatuschecker.Check(ctx, xurl.HTTP("localhost:1317")+"/node_info")
		if err == nil && !ok {
			err = errors.New("app is not online")
		}
		return err
	}
	return backoff.Retry(checkAlive, backoff.WithContext(backoff.NewConstantBackOff(time.Second), ctx))
}

// Must fails the immediately if not ok.
// t.Fail() needs to be called for the failing tests before running Must().
func (e env) Must(ok bool) {
	if !ok {
		e.t.FailNow()
	}
}

// Home returns user's home dir.
func (e env) Home() string {
	home, err := os.UserHomeDir()
	require.NoError(e.t, err)
	return home
}
