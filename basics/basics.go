package basics

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// HexStringToBase64String takes a hex string and returns a base64 encoded string.
// If the input string is not a valid hex string, an error is returned.
func HexStringToBase64String(hexString string) (string, error) {
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
}

// XOR takes two byte slices and returns a new byte slice that is the result of
// XORing the two input slices. If the input slices are not the same length, an
// error is returned.
func XOR(buffer1, buffer2 []byte) ([]byte, error) {
	if len(buffer1) != len(buffer2) {
		return nil, fmt.Errorf("buffers must have equal length")
	}

	result := make([]byte, len(buffer1))
	for i := 0; i < len(buffer1); i++ {
		result[i] = buffer1[i] ^ buffer2[i]
	}

	return result, nil
}
