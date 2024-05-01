package basics

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
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

var letterFrequencies = map[string]float64{
	"A": 8.167,
	"B": 1.492,
	"C": 2.782,
	"D": 4.253,
	"E": 12.702,
	"F": 2.228,
	"G": 2.015,
	"H": 6.094,
	"I": 6.966,
	"J": 0.153,
	"K": 0.772,
	"L": 4.025,
	"M": 2.406,
	"N": 6.749,
	"O": 7.507,
	"P": 1.929,
	"Q": 0.095,
	"R": 5.987,
	"S": 6.327,
	"T": 9.056,
	"U": 2.758,
	"V": 0.978,
	"W": 2.360,
	"X": 0.150,
	"Y": 1.974,
	"Z": 0.074,
}

func singleByteXOR(buffer []byte, key byte) []byte {
	result := make([]byte, len(buffer))
	for i := range len(buffer) {
		result[i] = buffer[i] ^ key
	}

	return result
}

var text = regexp.MustCompile("^[a-zA-Z ]$")

func isAlphabetic(str string) bool {
	return text.MatchString(str)
}

func calculateScore(buffer []byte) float64 {
	score := 0.0
	for _, b := range buffer {
		str := string(b)
		if isAlphabetic(str) {
			score += letterFrequencies[strings.ToUpper(str)]
		} else {
			score -= 10.0
		}
	}

	return score
}

func DecryptTheMessage(hexString string) ([]byte, error) {
	cipher, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, err
	}

	bestKey := byte(0)
	bestScore := 0.0
	for k := range 256 {
		key := byte(k)
		decrypted := singleByteXOR(cipher, key)

		score := calculateScore(decrypted)

		if score > bestScore {
			bestScore = score
			bestKey = key
		}
	}

	decryptedMessage := singleByteXOR(cipher, bestKey)
	return decryptedMessage, nil
}
