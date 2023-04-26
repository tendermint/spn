package encoding

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/tendermint/spn/cmd"
	"github.com/tendermint/spn/testutil/sample"
)

// MakeTestEncodingConfig creates a test EncodingConfig for an amino based test configuration.
func MakeTestEncodingConfig() cmd.EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := sample.InterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(marshaler, tx.DefaultSignModes)

	std.RegisterLegacyAminoCodec(amino)
	std.RegisterInterfaces(interfaceRegistry)

	return cmd.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             amino,
	}
}
