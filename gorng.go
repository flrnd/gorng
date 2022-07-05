/*
 * Check: https://blog.gopheracademy.com/advent-2017/a-tale-of-two-rands/
 */

package gorng

import (
	"crypto/rand"
	"encoding/binary"
	"math/big"
	mrand "math/rand"
)

var Rng = randInit()

type cryptoSource struct{}

func (s *cryptoSource) Seed(seed int64) {}

// Uint64 returns a securely generated int64 value.
func (s *cryptoSource) Uint64() (value uint64) {
	binary.Read(rand.Reader, binary.BigEndian, &value)
	return value
}

func (s *cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func randInit() (myRand *mrand.Rand) {
	var src mrand.Source64 = &cryptoSource{}
	myRand = mrand.New(src)

	return myRand
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if fails
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GenerateRandomInt returns securely generated random integer.
// It will return an error if fails
func GenerateRandomInt(n int64) (int, error) {
	num, err := rand.Int(rand.Reader, big.NewInt(n))
	if err != nil {
		return -1, err
	}
	return int(num.Int64()), nil
}
