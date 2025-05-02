package address_curve

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
