package block

func PKCS7Padding(data []byte, blockSize int) ([]byte, error) {
	rest := (len(data) % blockSize)
	if rest == 0 {
		return data, nil
	}

	paddingSize := blockSize - rest
	padded := make([]byte, len(data)+paddingSize)

	copy(padded, data)

	for i := len(data); i < len(padded); i++ {
		padded[i] = byte(paddingSize)
	}

	return padded, nil
}
