package address_curve

import (
	"crypto/sha3"

	"github.com/btcsuite/btcutil/base58"
)

// PublicKeyToAddress converts a public key to an address.
func PublicKeyToAddress(pubKey []byte) string {
	hash := sha3.Sum256(pubKey)
	addressBody := base58.Encode(hash[:])
	return addressBody
}
