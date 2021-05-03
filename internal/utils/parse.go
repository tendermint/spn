package utils

import "strconv"

// ParseInt32 convert string in base 10 to 32 bit signed integer.
// For the details see strconv.ParseInt.
func ParseInt32(s string) (int32, error) {
	rst, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(rst), nil
}
