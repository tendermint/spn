package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	profile "github.com/tendermint/spn/x/profile/types"
	"github.com/tendermint/spn/x/project/types"
)

func TestMsgEditProject_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgEditProject
		err  error
	}{
		{
			name: "should allow validation of msg with both name and metadata modified",
			msg: types.MsgEditProject{
				ProjectID:   0,
				Coordinator: sample.Address(r),
				Name:        sample.ProjectName(r),
				Metadata:    sample.Metadata(r, 20),
			},
		},
		{
			name: "should allow validation of msg with name modified",
			msg: types.MsgEditProject{
				ProjectID:   0,
				Coordinator: sample.Address(r),
				Name:        sample.ProjectName(r),
				Metadata:    []byte{},
			},
		},
		{
			name: "should allow validation of msg with metadata modified",
			msg: types.MsgEditProject{
				ProjectID:   0,
				Coordinator: sample.Address(r),
				Name:        "",
				Metadata:    sample.Metadata(r, 20),
			},
		},
		{
			name: "should prevent validation of msg with invalid project name",
			msg: types.MsgEditProject{
				ProjectID:   0,
				Coordinator: sample.Address(r),
				Name:        invalidProjectName,
				Metadata:    sample.Metadata(r, 20),
			},
			err: types.ErrInvalidProjectName,
		},
		{
			name: "should prevent validation of msg with invalid coordinator address",
			msg: types.MsgEditProject{
				ProjectID:   0,
				Coordinator: "invalid_address",
				Name:        sample.ProjectName(r),
				Metadata:    sample.Metadata(r, 20),
			},
			err: profile.ErrInvalidCoordAddress,
		},
		{
			name: "should prevent validation of msg with no fields modified",
			msg: types.MsgEditProject{
				ProjectID:   0,
				Coordinator: sample.Address(r),
				Name:        "",
				Metadata:    []byte{},
			},
			err: types.ErrCannotUpdateProject,
		},
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
