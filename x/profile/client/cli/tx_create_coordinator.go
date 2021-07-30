package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/profile/types"
)

const (
	flagIdentity = "identity"
	flagWebsite  = "website"
	flagDetails  = "details"
)

func CmdCreateCoordinator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-coordinator",
		Short: "Create a new coordinator profile",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			identity, err := cmd.Flags().GetString(flagIdentity)
			if err != nil {
				return err
			}

			website, err := cmd.Flags().GetString(flagWebsite)
			if err != nil {
				return err
			}

			details, err := cmd.Flags().GetString(flagDetails)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateCoordinator(
				clientCtx.GetFromAddress().String(),
				identity,
				website,
				details,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagIdentity, "", "coordinator identity")
	cmd.Flags().String(flagWebsite, "", "coordinator website url")
	cmd.Flags().String(flagDetails, "", "coordinator details")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
