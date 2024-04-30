package basics_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/stefanalfbo/cryptopals/basics"
)

func TestXOR(t *testing.T) {
	buffer1, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	buffer2, _ := hex.DecodeString("686974207468652062756c6c277320657965")
	expectedResult, _ := hex.DecodeString("746865206b696420646f6e277420706c6179")

	result, err := basics.XOR(buffer1, buffer2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !bytes.Equal(result, expectedResult) {
		t.Errorf("Expected %x, but got %x", expectedResult, result)
	}
}
