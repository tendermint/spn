package utils

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseInt32(t *testing.T) {
	for _, tc := range []struct {
		description string
		input       string
		result      int32
		err         string
	}{
		{
			description: "regular",
			input:       "10",
			result:      10,
		},
		{
			description: "overflow",
			input:       strconv.Itoa(math.MaxInt64),
			err:         "strconv.ParseInt: parsing \"%s\": value out of range",
		},
		{
			description: "not-integer",
			input:       "teepot",
			err:         "strconv.ParseInt: parsing \"%s\": invalid syntax",
		},
	} {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			rst, err := ParseInt32(tc.input)
			if err != nil {
				if tc.err != "" {
					require.EqualError(t, err, fmt.Sprintf(tc.err, tc.input))
				} else {
					require.NoError(t, err)
				}
			}
			require.Equal(t, tc.result, rst)
		})
	}
}
