package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	codec "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/launch/types"
)

var _ = strconv.Itoa(0)

func CmdRequestAddVestedAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-add-vested-account [chain-id] [starting-balance] [vesting-coins] [vesting-end-time]",
		Short: "Request to add a vested account",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			startingBalance, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return fmt.Errorf("failed to parse coins: %w", err)
			}

			vestingCoins, err := sdk.ParseCoinsNormalized(args[2])
			if err != nil {
				return fmt.Errorf("failed to parse coins: %w", err)
			}

			endTime, _ := strconv.ParseUint(args[3], 10, 64)

			delayedVesting, err := codec.NewAnyWithValue(&types.DelayedVesting{
				Vesting: vestingCoins,
				EndTime: int64(endTime),
			})
			if err != nil {
				return err
			}

			msg := types.NewMsgRequestAddVestedAccount(
				clientCtx.GetFromAddress().String(),
				args[0],
				startingBalance,
				delayedVesting,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
