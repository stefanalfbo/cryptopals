package basics_test

import (
	"testing"

	"github.com/stefanalfbo/cryptopals/basics"
)

func TestDetectAESInECBMode(t *testing.T) {
	plaintext := []byte("YELLOW SUBMARINEYELLOW SUBMARINE")
	key := []byte("YELLOW SUBMARINE")

	ciphertext, _ := basics.EncryptAESWithECBMode(plaintext, key)
	keySize := len(key)

	if !basics.DetectAESInECBMode(ciphertext, keySize) {
		t.Errorf("Expected to detect ECB mode")
	}
}

func TestDetectAESInECBModeFile(t *testing.T) {
	keySize := 16

	candidate, _ := basics.DetectAESInECBModeFile("8.txt", keySize)

	if candidate == nil {
		t.Errorf("Expected to find a candidate")
	}
}
