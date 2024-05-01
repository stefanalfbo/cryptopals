package basics_test

import (
	"bytes"
	"testing"

	"github.com/stefanalfbo/cryptopals/basics"
)

func TestFindEncryptedLine(t *testing.T) {
	bestLine, err := basics.FindEncryptedLine("4.txt")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	expected := []byte("Now that the party is jumping\n")
	if !bytes.Equal(bestLine, expected) {
		t.Errorf("Expected: %v, Got: %v", "Now that the party is jumping\n", bestLine)
	}
}
