package cli_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/tendermint/spn/testutil/networksuite"
)

// QueryTestSuite is a test suite for query tests
type QueryTestSuite struct {
	networksuite.NetworkTestSuite
}

// TestQueryTestSuite runs test of the query suite
func TestQueryTestSuite(t *testing.T) {
	suite.Run(t, new(QueryTestSuite))
}
