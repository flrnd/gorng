package gorng

import (
	crand "crypto/rand"
	"errors"
	"math/big"
)

// --- Secure Helpers ---
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := crand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GenerateRandomInt(n int64) (int, error) {
	if n <= 0 {
		return -1, errors.New("n must be > 0")
	}
	num, err := crand.Int(crand.Reader, big.NewInt(n))
	if err != nil {
		return -1, err
	}
	return int(num.Int64()), nil
}
