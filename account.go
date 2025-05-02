package address_curve

import (
	"fmt"

	"github.com/btcsuite/btcutil/base58"
)

type Account struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

// CreateAccount generates a new account with a private key, public key, and address.
// It uses the Ed25519 algorithm to generate a private key and derives the public key and address from it.
func CreateAccount() (*Account, error) {
	privKey, err := generatePrivateKey()
	if err != nil {
		return nil, err
	}
	leaves := deriveLeafNodes(privKey)
	merkleRoot := buildMerkleTree(leaves)

	return &Account{
		PrivateKey: EncodePrivateKey(privKey),
		PublicKey:  fmt.Sprintf("%x", merkleRoot),
		Address:    PublicKeyToAddress(merkleRoot),
	}, nil
}

// ComputeAddress computes the account from a given private key.
func ComputeAddress(encodedPrivKey string) (*Account, error) {
	privKey := base58.Decode(encodedPrivKey)
	if len(privKey) != 64 {
		return nil, fmt.Errorf("invalid private key length: expected 64 bytes, got %d", len(privKey))
	}

	leaves := deriveLeafNodes(privKey)

	merkleRoot := buildMerkleTree(leaves)

	publicKey := PublicKeyToAddress(merkleRoot)

	return &Account{
		PrivateKey: encodedPrivKey,
		PublicKey:  fmt.Sprintf("%x", merkleRoot),
		Address:    publicKey,
	}, nil
}
