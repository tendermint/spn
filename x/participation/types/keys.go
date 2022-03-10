package types

import (
	"strconv"
)

const (
	// ModuleName defines the module name
	ModuleName = "participation"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_participation"

	// UsedAllocationsKeyPrefix is the prefix to retrieve all UsedAllocations
	UsedAllocationsKeyPrefix = "UsedAllocations/value/"

	// AuctionUsedAllocationsKeyPrefix is the prefix to retrieve all AuctionUsedAllocations
	AuctionUsedAllocationsKeyPrefix = "AuctionUsedAllocations/value/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// UsedAllocationsKey returns the store key to retrieve a UsedAllocations from the address field
func UsedAllocationsKey(address string) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}

// AuctionUsedAllocationsKey returns the store key to retrieve a AuctionUsedAllocations from the address and auctionID fields
func AuctionUsedAllocationsKey(address string, auctionID uint64) []byte {
	var key []byte

	addressBytes := []byte(address)
	auctionIDBytes := []byte(strconv.FormatUint(auctionID, 10))
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)
	key = append(key, auctionIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
