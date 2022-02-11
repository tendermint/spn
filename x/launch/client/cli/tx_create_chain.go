package cli

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/launch/types"
)

const (
	flagGenesisURL = "genesis-url"
	flagCampaignID = "campaign-id"
)

func CmdCreateChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-chain [genesis-chain-id] [source-url] [source-hash]",
		Short: "Create a new chain for launch",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argCampaignID, err := cmd.Flags().GetInt64(flagCampaignID)
			if err != nil {
				return err
			}

			hasCampaign := false
			campaignID := uint64(0)
			if argCampaignID >= 0 {
				hasCampaign = true
				campaignID = uint64(argCampaignID)
			}

			genesisURL, err := cmd.Flags().GetString(flagGenesisURL)
			if err != nil {
				return err
			}
			var genesisHash string
			if genesisURL != "" {
				genesisHash, err = getHashFromURL(cmd.Context(), genesisURL)
				if err != nil {
					return err
				}
			}

			msg := types.NewMsgCreateChain(
				clientCtx.GetFromAddress().String(),
				args[0],
				args[1],
				args[2],
				genesisURL,
				genesisHash,
				hasCampaign,
				campaignID,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagGenesisURL, "", "URL for a custom genesis")
	cmd.Flags().Int64(flagCampaignID, -1, "The campaign id")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// getHashFromURL fetches content from url and returns the hash based on the genesis hash method
func getHashFromURL(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("genesis url fetch error %s", res.Status)
	}
	initialGenesis, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return types.GenesisURLHash(string(initialGenesis)), nil
}
