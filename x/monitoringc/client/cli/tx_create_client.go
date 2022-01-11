package cli

import (
	"encoding/base64"
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	committypes "github.com/cosmos/ibc-go/modules/core/23-commitment/types"
	ibctmtypes "github.com/cosmos/ibc-go/modules/light-clients/07-tendermint/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/monitoringc/types"
	"github.com/tendermint/tendermint/light"
	"time"
)

func CmdCreateClient() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-client [launch-id]",
		Short: "Create a verified client ID to connect to the chain with the specified launch ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			launchID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// sample client State
			chainID := "orbit-1"
			unbondingPeriod := time.Hour*24*21
			trustingPeriod := unbondingPeriod-1
			latestHeight := clienttypes.NewHeight(1, 1)
			clientState := ibctmtypes.NewClientState(
				chainID,
				ibctmtypes.NewFractionFromTm(light.DefaultTrustLevel),
				trustingPeriod,
				unbondingPeriod,
				time.Minute*10,
				latestHeight,
				committypes.GetSDKSpecs(),
				[]string{"upgrade", "upgradedIBCState"},
				true,
				true,
			)

			timestamp, err := time.Parse(time.RFC3339, "2022-01-11T08:25:36.020826Z")
			if err != nil {
				return err
			}
			nextValSetHash, _ := hex.DecodeString("78FD31D1FA7E68C2527729505D9B47A80A1960B3E95A8C47A4FE2F13FCDD9731")
			rootHash, _ := base64.StdEncoding.DecodeString("47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=")
			consensusState := ibctmtypes.NewConsensusState(
				timestamp,
				committypes.NewMerkleRoot(rootHash),
				nextValSetHash,
			)

			msg := types.NewMsgCreateClient(
				clientCtx.GetFromAddress().String(),
				launchID,
				*clientState,
				*consensusState,
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
