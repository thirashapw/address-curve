package address_curve

import (
	"crypto/sha3"

	"github.com/btcsuite/btcutil/base58"
)

func PublicKeyToAddress(pubKey []byte) string {
	hash := sha3.Sum256(pubKey)
	addressBody := base58.Encode(hash[:])
	return addressBody
}
