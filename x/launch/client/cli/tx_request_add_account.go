package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/launch/types"
)

const (
	flagAccountAddress = "account-address"
)

func CmdRequestAddAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-add-account [launch-id] [coins]",
		Short: "Request to add an account",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coins, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return fmt.Errorf("failed to parse coins: %w", err)
			}

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
				types.NewGenesisAccount(launchID, accountAddr, coins),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagAccountAddress, "", "Address of the genesis account to request")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
