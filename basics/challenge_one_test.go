package basics_test

import (
	"testing"

	"github.com/stefanalfbo/cryptopals/basics"
)

func TestConvertToBase64(t *testing.T) {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expectedBase64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	result, err := basics.HexStringToBase64String(hexString)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != expectedBase64 {
		t.Errorf("Expected base64: %s, but got: %s", expectedBase64, result)
	}
}
