package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MonitoringHistoryKeyPrefix is the prefix to retrieve all MonitoringHistory
	MonitoringHistoryKeyPrefix = "MonitoringHistory/value/"
)

// MonitoringHistoryKey returns the store key to retrieve a MonitoringHistory from the index fields
func MonitoringHistoryKey(launchID uint64) []byte {
	return append(spntypes.UintBytes(launchID), byte('/'))
}
