package basics_test

import (
	"testing"

	"github.com/stefanalfbo/cryptopals/basics"
)

func TestDecryptAESWithECBMode(t *testing.T) {
	plaintext, err := basics.DecryptAESWithECBModeFile("7.txt", []byte("YELLOW SUBMARINE"))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	t.Logf("Plaintext: %s", plaintext)
}

func TestEncryptAESWithECBMode(t *testing.T) {
	plaintext, err := basics.DecryptAESWithECBModeFile("7.txt", []byte("YELLOW SUBMARINE"))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	ciphertext, err := basics.EncryptAESWithECBMode(plaintext, []byte("YELLOW SUBMARINE"))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	decryptedPlaintext, err := basics.DecryptAESWithECBMode(ciphertext, []byte("YELLOW SUBMARINE"))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if string(decryptedPlaintext) != string(plaintext) {
		t.Errorf("Expected plaintext: %s, got: %s", string(plaintext), string(decryptedPlaintext))
	}
}
