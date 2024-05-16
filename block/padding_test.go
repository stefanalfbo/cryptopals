package block_test

import (
	"testing"

	"github.com/stefanalfbo/cryptopals/block"
)

func TestPKCS7Padding(t *testing.T) {
	data := []byte("YELLOW SUBMARINE")

	t.Run("Padding needed", func(t *testing.T) {
		expected := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")
		blockSize := 20

		actual, err := block.PKCS7Padding(data, blockSize)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if string(actual) != string(expected) {
			t.Errorf("Expected %v, was %v", expected, actual)
		}
	})

	t.Run("No padding needed", func(t *testing.T) {
		blockSize := 16
		actual, err := block.PKCS7Padding(data, blockSize)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if string(actual) != string(data) {
			t.Errorf("Expected %v, was %v", data, actual)
		}
	})

	t.Run("Validate that the value for the padding adapts to the length needed", func(t *testing.T) {
		expected := []byte("YELLOW SUBMARINE\x06\x06\x06\x06\x06\x06")
		blockSize := 22

		actual, err := block.PKCS7Padding(data, blockSize)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if string(actual) != string(expected) {
			t.Errorf("Expected %v, was %v", expected, actual)
		}
	})
}
