# gorng — Go Random Number Generator Module

gorng is a lightweight Go module for generating both cryptographically secure and deterministic pseudo-random numbers, with pluggable backends using Go 1.22’s new math/rand/v2 standard library package.

It supports:
  - 🔐 Secure random bytes and integers (crypto/rand)
  - 🎲 Deterministic PRNGs:
    - ChaCha8 — high-quality stream cipher for reproducible randomness
    - PCG — small, fast, statistically sound

## Features

- Simple interface: PRNG (compatible with *randv2.Rand)
- Secure helpers: GenerateRandomBytes, GenerateRandomInt
- Supports:
  - NewChaCha8PRNG(seed [32]byte)
  - NewChaCha8SeededPRNG() — secure seeded
  - NewPCGPRNG(seed, stream uint64)
  - NewPCGRandomPRNG() — securely seeded
  - NewCryptoPRNG() — cryptographically secure

### ⚠️ Breaking Changes in v0.0.3

If you're upgrading from v0.0.2, please note:
- ❌ Removed the previous math/rand.Source64-based crypto implementation
- ✅ PRNG is now an alias for *randv2.Rand
- 🆕 ChaCha8 and PCG must now be wrapped with randv2.New(...) as required by Go 1.22’s rand/v2 model
- ✅ cryptoPRNG remains available and implements the same interface manually

### 📦 Requirements

Go 1.22+
