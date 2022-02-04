package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MonitoringHistoryKeyPrefix is the prefix to retrieve all MonitoringHistory
	MonitoringHistoryKeyPrefix = "MonitoringHistory/value/"
)

// MonitoringHistoryKey returns the store key to retrieve a MonitoringHistory from the index fields
func MonitoringHistoryKey(
	launchID uint64,
) []byte {
	var key []byte

	launchIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(launchIDBytes, launchID)
	key = append(key, launchIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
