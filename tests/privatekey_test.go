package tests

import (
	"testing"

	address_curve "github.com/thirashapw/address-curve"
)

func TestPrivateKey(t *testing.T) {
	privKey, err := address_curve.GeneratePrivateKey()
	if err != nil {
		t.Fatalf("failed to generate private key: %v", err)
	}
	t.Logf("private key generated successfully: %x", privKey)

}
