package cli

import (
	"io/ioutil"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/launch/types"
)

const (
	flagValidatorAddress = "validator-address"
)

func CmdRequestAddValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-add-validator [launch-id] [gentx-file] [consensus-public-key] [self-delegation] [peer]",
		Short: "Send a request for a genesis validator",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Read self-delegation
			selfDelegation, err := sdk.ParseCoinNormalized(args[3])
			if err != nil {
				return err
			}

			// Read gentxFile
			gentxBytes, err := ioutil.ReadFile(args[1])
			if err != nil {
				return err
			}

			launchID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			fromAddr := clientCtx.GetFromAddress().String()
			valAddr, _ := cmd.Flags().GetString(flagValidatorAddress)
			if valAddr == "" {
				valAddr = fromAddr
			}

			msg := types.NewMsgRequestAddValidator(
				clientCtx.GetFromAddress().String(),
				launchID,
				valAddr,
				gentxBytes,
				[]byte(args[2]),
				selfDelegation,
				args[4],
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagValidatorAddress, "", "Address of the genesis validator to request")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
