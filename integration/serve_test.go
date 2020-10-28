package integration_test

import (
	"context"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServeApp(t *testing.T) {
	t.Parallel()

	projectDir, _ := os.Getwd()
	projectDir = path.Join(projectDir, "..")

	var (
		env = newEnv(t)
	)

	var (
		ctx, cancel       = context.WithTimeout(env.Ctx(), serveTimeout)
		isBackendAliveErr error
	)
	go func() {
		defer cancel()
		isBackendAliveErr = env.IsAppServed(ctx)
	}()
	env.Must(env.Serve("should serve", projectDir, ExecCtx(ctx)))

	require.NoError(t, isBackendAliveErr, "app cannot get online in time")
}
