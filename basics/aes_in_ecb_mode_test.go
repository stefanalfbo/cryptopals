package basics_test

import (
	"testing"

	"github.com/stefanalfbo/cryptopals/basics"
)

func TestDecryptAESWithECBMode(t *testing.T) {
	plaintext, err := basics.DecryptAESWithECBMode("7.txt", []byte("YELLOW SUBMARINE"))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	t.Logf("Plaintext: %s", plaintext)
}
