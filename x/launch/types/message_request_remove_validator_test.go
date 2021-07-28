package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMsgRequestRemoveValidator_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRequestRemoveValidator
		err  error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
