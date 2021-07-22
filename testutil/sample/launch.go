package sample

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	launch "github.com/tendermint/spn/x/launch/types"
	"time"
)

func GenesisAccount(chainID, address string) *launch.GenesisAccount {
	return &launch.GenesisAccount{
		ChainID: chainID,
		Address: address,
		Coins:   Coins(),
	}
}

func Request(chainID string) *launch.Request {
	content, err := types.NewAnyWithValue(GenesisAccount(chainID, AccAddress()))
	if err != nil {
		panic(err)
	}

	return &launch.Request{
		ChainID: chainID,
		Creator: AccAddress(),
		CreatedAt: time.Now().Unix(),
		Content: content,
	}
}