package address_curve

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
)

// GeneratePrivateKey generates a new Ed25519 private key.
// It returns the private key as a byte slice and an error if any occurs.
// The private key is 64 bytes long, which is the standard size for Ed25519 keys.
func generatePrivateKey() (ed25519.PrivateKey, error) {
	privKey := make([]byte, 64)
	_, err := rand.Read(privKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %w", err)
	}
	return privKey, nil
}

func EncodePrivateKey(privKey []byte) string {
	return base58.Encode(privKey)
}
