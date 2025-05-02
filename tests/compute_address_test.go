package tests

import (
	"testing"

	address_curve "github.com/thirashapw/address-curve"
)

func TestFromPrivateKey(t *testing.T) {
	account, err := address_curve.CreateAccount()
	if err != nil {
		t.Fatalf("failed to generate private key: %v", err)
	}
	
	t.Logf("private key generated successfully: %s", account.PrivateKey)
	t.Logf("public key generated successfully: %s", account.PublicKey)
	t.Logf("address generated successfully: %s", account.Address)

	loaded_account, err := address_curve.ComputeAddress(account.PrivateKey)
	if err != nil {
		t.Fatalf("failed to generate private key: %v", err)
	}
	t.Logf("private key loaded successfully: %s", loaded_account.PrivateKey)
	t.Logf("public key loaded successfully: %s", loaded_account.PublicKey)
	t.Logf("address loaded successfully: %s", loaded_account.Address)
}
