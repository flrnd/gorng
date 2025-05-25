// math/rand/v2: revised API for math/rand
// https://github.com/golang/go/issues/61716

package gorng

import (
	crand "crypto/rand"
	"encoding/binary"
	randv2 "math/rand/v2"
)

// --- PRNG Interface ---
type PRNG = *randv2.Rand

// --- ChaCha8 PRNG ---
func NewChaCha8PRNG(seed [32]byte) PRNG {
	return randv2.New(randv2.NewChaCha8(seed))
}

func NewChaCha8SeededPRNG() (PRNG, error) {
	b, err := GenerateRandomBytes(32)
	if err != nil {
		return nil, err
	}

	var seed [32]byte

	copy(seed[:], b)

	return NewChaCha8PRNG(seed), nil
}

// --- PCG PRNG ---
func NewPCGPRNG(seed, stream uint64) PRNG {
	return randv2.New(randv2.NewPCG(seed, stream))
}

func NewPCGRandomPRNG() (PRNG, error) {
	var seeds [2]uint64

	err := binary.Read(crand.Reader, binary.BigEndian, &seeds)
	if err != nil {
		return nil, err
	}

	return NewPCGPRNG(seeds[0], seeds[1]), nil
}

// --- Crypto PRNG ---
type cryptoPRNG struct{}

func (c cryptoPRNG) IntN(n int) int {
	val, err := GenerateRandomInt(int64(n))

	if err != nil {
		panic("crypto PRNG failed")
	}

	return val
}

func (c cryptoPRNG) Uint64() uint64 {
	var val uint64

	_ = binary.Read(crand.Reader, binary.BigEndian, &val)

	return val
}

func NewCryptoPRNG() cryptoPRNG {
	return cryptoPRNG{}
}
