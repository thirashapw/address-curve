package address_curve

import (
	"fmt"

	"github.com/btcsuite/btcutil/base58"
)

type Account struct {
	PrivateKey string
	PublicKey  string
}

func CreateAccount() (*Account, error) {
	privKey, err := GeneratePrivateKey()
	if err != nil {
		return nil, err
	}
	leaves := deriveLeafNodes(privKey)
	merkleRoot := buildMerkleTree(leaves)

	return &Account{
		PrivateKey: EncodePrivateKey(privKey),
		PublicKey:  PublicKeyToAddress(merkleRoot),
	}, nil
}

func ComputeAddress(encodedPrivKey string) (*Account, error) {
	privKey := base58.Decode(encodedPrivKey)
	if len(privKey) != 60 {
		return nil, fmt.Errorf("invalid private key length: expected 60 bytes, got %d", len(privKey))
	}

	leaves := deriveLeafNodes(privKey)

	merkleRoot := buildMerkleTree(leaves)

	publicKey := PublicKeyToAddress(merkleRoot)

	return &Account{
		PrivateKey: encodedPrivKey,
		PublicKey:  publicKey,
	}, nil
}
