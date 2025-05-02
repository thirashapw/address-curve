package tests

import (
	"testing"

	address_curve "github.com/thirashapw/address-curve"
)

func TestFromPrivateKey(t *testing.T) {
	account, err := address_curve.ComputeAddress("8FnBwL2YGEvKsLyzNBUiSfkVZMiWPr4LhVpyvE5pbCBjWE72CUnoEA2aYqx5fcsQvdCcUqhSvJKFfytA7T")
	if err != nil {
		t.Fatalf("failed to generate private key: %v", err)
	}
	t.Logf("private key loaded successfully: %s", account.PrivateKey)
	t.Logf("public key loaded successfully: %s", account.PublicKey)
}
