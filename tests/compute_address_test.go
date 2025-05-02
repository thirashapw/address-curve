package tests

import (
	"testing"

	address_curve "github.com/thirashapw/address-curve"
)

func TestFromPrivateKey(t *testing.T) {
	account, err := address_curve.ComputeAddress("4YyU853RCrsnexE5TvyQhzsUPWey8g1vGBWMx83wnCAgZbqmLRwVKZi1A8C6HZ4z7vkT6KdzTksEdqERe2CrN5Wa")
	if err != nil {
		t.Fatalf("failed to generate private key: %v", err)
	}
	t.Logf("private key loaded successfully: %s", account.PrivateKey)
	t.Logf("public key loaded successfully: %s", account.PublicKey)
	t.Logf("address loaded successfully: %s", account.Address)
}