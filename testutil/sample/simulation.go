package sample

import (
	"math/rand"
	"time"
)

// Rand returns a sample Rand object for randomness
func Rand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().Unix()))
}