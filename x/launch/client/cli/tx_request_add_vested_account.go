package cli

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	codec "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/launch/types"
)

func CmdRequestAddVestedAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-add-vested-account [chainID] [coins] [options]",
		Short: "Request to add a vested account",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coins, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return fmt.Errorf("failed to parse coins: %w", err)
			}

			delayedVesting, err := codec.NewAnyWithValue(&types.DelayedVesting{
				Vesting: coins,
				EndTime: time.Now().Unix(),
			})
			if err != nil {
				return err
			}

			// TODO: return an error if the launch is triggered for the chain

			msg := types.NewMsgRequestAddVestedAccount(
				clientCtx.GetFromAddress().String(),
				args[0],
				coins,
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
