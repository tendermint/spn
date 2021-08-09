package cli

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseList(t *testing.T) {
	cases := []struct {
		list   string
		parsed []uint64
	}{
		{"1,2,3", []uint64{1, 2, 3}},
		{"1, 2,3 ", []uint64{1, 2, 3}},
		{",1, 2,", []uint64{1, 2}},
		{"1/3 ", []uint64{1, 2, 3}},
		{"1/3,8 ", []uint64{1, 2, 3, 8}},
		{"1/3,8/11 ", []uint64{1, 2, 3, 8, 9, 10, 11}},
		{"1/3,8/11,33 ", []uint64{1, 2, 3, 8, 9, 10, 11, 33}},
		{"1/3,8/11,33/36 ", []uint64{1, 2, 3, 8, 9, 10, 11, 33, 34, 35, 36}},
		{",", []uint64{}},
		{",/", []uint64{}},
		{",10/", []uint64{10}},
		{"10/", []uint64{10}},
		{"/10", []uint64{10}},
		{"10/10", []uint64{10}},
	}
	for _, tt := range cases {
		t.Run("list "+tt.list, func(t *testing.T) {
			parsed, err := parseList(tt.list)
			require.NoError(t, err)
			require.Equal(t, tt.parsed, parsed)
		})
	}
}
