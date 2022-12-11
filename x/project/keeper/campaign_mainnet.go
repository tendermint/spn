package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// IsProjectMainnetLaunchTriggered returns true if the provided project has an associated mainnet chain whose launch
// has been already triggered
func (k Keeper) IsProjectMainnetLaunchTriggered(ctx sdk.Context, projectID uint64) (bool, error) {
	project, found := k.GetProject(ctx, projectID)
	if !found {
		return false, fmt.Errorf("project %d not found", projectID)
	}

	if project.MainnetInitialized {
		chain, found := k.launchKeeper.GetChain(ctx, project.MainnetID)
		if !found {
			return false, fmt.Errorf("mainnet chain %d for project %d not found", project.MainnetID, projectID)
		}
		if !chain.IsMainnet {
			return false, fmt.Errorf("chain %d for project %d is not a mainnet chain", project.MainnetID, projectID)
		}
		if chain.LaunchTriggered {
			return true, nil
		}
	}
	return false, nil
}
