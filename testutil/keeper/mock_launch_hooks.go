package keeper

import (
	launchtypes "github.com/tendermint/spn/x/launch/types"
)

//go:generate mockery --name LaunchHooks --filename mock_launch_hooks.go --case underscore --output ./mocks
type LaunchHooks interface {
	launchtypes.LaunchHooks
}
