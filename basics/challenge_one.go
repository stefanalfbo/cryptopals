package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func convertToBase64(hexString string) (string, error) {
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
}

func main() {
	theString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d" // Replace with your hex string

	str, err := convertToBase64(theString)
	if err != nil {
		fmt.Println("Error decoding hexadecimal string:", err)
		return
	}

	fmt.Println("Base64:", str)
}
