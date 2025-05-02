package tests

import (
	"testing"

	address_curve "github.com/thirashapw/address-curve"
)

func TestNewAccount(t *testing.T) {
	account, err := address_curve.CreateAccount()
	if err != nil {
		t.Fatalf("failed to generate private key: %v", err)
	}
	t.Logf("private key fresh generated successfully: %s", account.PrivateKey)
	t.Logf("public key fresh generated successfully: %s", account.PublicKey)
	t.Logf("address fresh generated successfully: %s", account.Address)
}
