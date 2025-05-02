package address_curve

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
)

func GeneratePrivateKey() (ed25519.PrivateKey, error) {
	privKey := make([]byte, 60)
	_, err := rand.Read(privKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %w", err)
	}
	return privKey, nil
}

func EncodePrivateKey(privKey []byte) string {
	return base58.Encode(privKey)
}
