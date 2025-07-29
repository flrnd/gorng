// Package gorng provides cryptographically secure random number
// generation by directly using the "crypto/rand" package, which is the
// standard and recommended approach in Go for security-sensitive applications.
package gorng

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// GenerateRandomBytes returns securely generated random bytes.
// It uses crypto/rand.Read, which obtains random data directly from the OS.
//
// n: The number of bytes to generate.
func GenerateRandomBytes(n int) ([]byte, error) {
	if n <= 0 {
		return nil, fmt.Errorf("number of bytes (n) must be a positive integer")
	}
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random bytes: %w", err)
	}
	return b, nil
}

// GenerateRandomInt64 returns a securely generated random integer in the range [0, max).
// It will return an error if max <= 0 or if the underlying crypto/rand.Int call fails.
// It uses crypto/rand.Int for a uniform random value, avoiding modulo bias.
//
// max: The upper bound for the random integer (exclusive).
func GenerateRandomInt64(max int64) (int64, error) {
	if max <= 0 {
		return 0, fmt.Errorf("max must be greater than 0")
	}

	num, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return 0, fmt.Errorf("failed to generate random int: %w", err)
	}

	return num.Int64(), nil
}
