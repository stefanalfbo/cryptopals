package basics_test

import (
	"bytes"
	"testing"

	"github.com/stefanalfbo/cryptopals/basics"
)

func TestDecryptTheMessage(t *testing.T) {
	hexString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	result, err := basics.DecryptTheMessage(hexString)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := []byte("Cooking MC's like a pound of bacon")
	if !bytes.Equal(result, expected) {
		t.Errorf("Expected message: %s, but got: %s", expected, result)
	}
}
