package basics

import (
	"crypto/aes"
)

func DecryptAESWithECBMode(filename string, key []byte) ([]byte, error) {
	ciphertext, err := decodeBase64File(filename)
	if err != nil {
		return nil, err
	}

	block, _ := aes.NewCipher(key)
	plaintext := make([]byte, len(ciphertext))
	blockSize := len(key)

	for start, end := 0, blockSize; start < len(ciphertext); start, end = start+blockSize, end+blockSize {
		block.Decrypt(plaintext[start:end], ciphertext[start:end])
	}

	return plaintext, nil
}
