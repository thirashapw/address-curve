package address_curve

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
)

func GeneratePrivateKey() (ed25519.PrivateKey, error) {
	privKey := make([]byte, 40)
	_, err := rand.Read(privKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %w", err)
	}
	return privKey, nil
}
