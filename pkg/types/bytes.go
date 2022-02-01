package types

import "encoding/binary"

// UintBytes convert uint64 to byte slice
func UintBytes(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
