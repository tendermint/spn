package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/profile/types"
)

const (
	flagMoniker         = "moniker"
	flagSecurityContact = "security-contact"
)

func CmdUpdateValidatorDescription() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-validator-description",
		Short: "Update a validator description",
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

			moniker, err := cmd.Flags().GetString(flagMoniker)
			if err != nil {
				return err
			}

			securityContact, err := cmd.Flags().GetString(flagSecurityContact)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateValidatorDescription(
				clientCtx.GetFromAddress().String(),
				identity,
				moniker,
				website,
				securityContact,
				details,
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
