package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/project/types"
)

const flagName = "flag-project-name"

func CmdEditProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-project [project-id]",
		Short: "Edit the project name or metadata",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			projectID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			name, err := cmd.Flags().GetString(flagName)
			if err != nil {
				return err
			}

			metadata, err := cmd.Flags().GetString(flagMetadata)
			if err != nil {
				return err
			}
			metadataBytes := []byte(metadata)

			msg := types.NewMsgEditProject(
				clientCtx.GetFromAddress().String(),
				projectID,
				name,
				metadataBytes,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagName, "", "Set name for the project")
	cmd.Flags().String(flagMetadata, "", "Set metadata field for the project")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
