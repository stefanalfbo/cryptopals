package basics

import (
	"crypto/aes"
)

func DecryptAESWithECBModeFile(filename string, key []byte) ([]byte, error) {
	ciphertext, err := decodeBase64File(filename)
	if err != nil {
		return nil, err
	}

	return DecryptAESWithECBMode(ciphertext, key)
}

func DecryptAESWithECBMode(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, len(ciphertext))
	blockSize := len(key)

	for start, end := 0, blockSize; start < len(ciphertext); start, end = start+blockSize, end+blockSize {
		block.Decrypt(plaintext[start:end], ciphertext[start:end])
	}

	return plaintext, nil
}

func EncryptAESWithECBModeFile(filename string, key []byte) ([]byte, error) {
	plaintext, err := decodeBase64File(filename)
	if err != nil {
		return nil, err
	}

	return EncryptAESWithECBMode(plaintext, key)
}

func EncryptAESWithECBMode(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, len(plaintext))
	blockSize := len(key)

	for start, end := 0, blockSize; start < len(plaintext); start, end = start+blockSize, end+blockSize {
		block.Encrypt(ciphertext[start:end], plaintext[start:end])
	}

	return ciphertext, nil
}
