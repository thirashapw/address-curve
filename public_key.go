package address_curve

import (
	"crypto/sha3"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
)

func PrivateKeyToPublicKey(privKey []byte) []byte {
	hash := sha3.Sum256(privKey)
	return hash[:25]
}

func EncodePublicKey(pubKey []byte, prefix string) string {
	if prefix == "" {
		return base58.Encode(pubKey)
	} else {
		return fmt.Sprintf("%s%s", prefix, base58.Encode(pubKey))
	}
}
