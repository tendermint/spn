package cli

import (
	"fmt"
	"strconv"

	"github.com/aws/smithy-go/time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/launch/types"
)

func CmdRequestAddVestingAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-add-vesting-account [launch-id] [total-balance] [vesting-coins] [vesting-end-time]",
		Short: "Request to add a vesting account",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			totalBalance, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return fmt.Errorf("failed to parse coins: %w", err)
			}

			vestingCoins, err := sdk.ParseCoinsNormalized(args[2])
			if err != nil {
				return fmt.Errorf("failed to parse coins: %w", err)
			}

			endTime, err := time.ParseDateTime(args[3])
			if err != nil {
				return err
			}

			delayedVesting := *types.NewDelayedVesting(totalBalance, vestingCoins, endTime)

			launchID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			fromAddr := clientCtx.GetFromAddress().String()
			accountAddr, _ := cmd.Flags().GetString(flagAccountAddress)
			if accountAddr == "" {
				accountAddr = fromAddr
			}

			msg := types.NewMsgSendRequest(
				fromAddr,
				launchID,
				types.NewVestingAccount(launchID, accountAddr, delayedVesting),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagAccountAddress, "", "Address of the vesting account to request")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
