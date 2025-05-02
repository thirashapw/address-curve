package tests

import (
	"testing"

	address_curve "github.com/thirashapw/address-curve"
)

func TestFromPrivateKey(t *testing.T) {
	account, err := address_curve.ComputeAddress("JDTkLBfvFK26KN7D54cvkre1qmiDGT8WZfMU4SVUn8oVmkyQ8rZQ5R46gYLYLTDjeRDm7HLzWD71Lm7DVQ")
	if err != nil {
		t.Fatalf("failed to generate private key: %v", err)
	}
	t.Logf("private key loaded successfully: %s", account.PrivateKey)
	t.Logf("public key loaded successfully: %s", account.PublicKey)
}
