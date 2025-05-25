# gorng

A minimal, extensible Go module for generating both cryptographically secure and deterministic pseudo-random numbers. Starting from v2, it uses Go 1.22’s modern `math/rand/v2` APIs with swappable PRNG backends like ChaCha8 and PCG.

---

## ✨ Features

- 🔐 **Cryptographically secure randomness** (`crypto/rand`)
- 🎲 **Deterministic PRNGs**:
  - **ChaCha8** — reproducible randomness with strong distribution
  - **PCG** — compact, fast, and statistically sound
- 🧩 Simple `PRNG` interface compatible with `*randv2.Rand`
- Designed for password generation, simulations, procedural content, and CLI utilities

---

## 📦 Versioning

This project uses Go modules with versioned import paths:

| Version | Import Path                        | Notes                                      |
|---------|------------------------------------|--------------------------------------------|
| v1.x    | `github.com/flrnd/gorng`    | Legacy version using `crypto/rand` and custom `rand.Source64` |
| v2.x    | `github.com/flrnd/gorng/v2` | ✅ Current version, based on `math/rand/v2` |

---

## ⚠️ Breaking Changes in v2

- `PRNG` is now an alias for `*randv2.Rand`, offering `IntN()` and `Uint64()` out of the box.
- ChaCha8 and PCG PRNGs are constructed via `randv2.New(...)`.
- Removed `math/rand.Source64`-based crypto RNG from v1.
- Requires **Go 1.22+**.

---

## ✅ Example Usage (v2)

```go
import "github.com/flrnd/gorng/v2"

func main() {
    // 🔐 Crypto PRNG
    crypto := gorng.NewCryptoPRNG()
    fmt.Println("Secure Int:", crypto.IntN(100))

    // 🌀 ChaCha8 (deterministic)
    var seed [32]byte
    copy(seed[:], "your 32-byte seed value here")
    chacha := gorng.NewChaCha8PRNG(seed)
    fmt.Println("ChaCha8:", chacha.IntN(100))

    // 🎲 PCG
    pcg := gorng.NewPCGPRNG(12345, 67890)
    fmt.Println("PCG:", pcg.IntN(100))
}
```