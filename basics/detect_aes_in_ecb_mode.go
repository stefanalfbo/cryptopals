package basics

import "encoding/hex"

func DetectAESInECBMode(ciphertext []byte, keySize int) bool {
	blocks := make(map[string]bool)

	for i := 0; i < len(ciphertext); i += keySize {
		block := string(ciphertext[i : i+keySize])
		if blocks[block] {
			return true
		}
		blocks[block] = true
	}

	return false
}

func DetectAESInECBModeFile(filename string, keySize int) ([]byte, error) {
	lines, err := readFileLines(filename)
	if err != nil {
		return nil, err
	}
	for _, line := range lines {
		decoded, err := hex.DecodeString(line)
		if err != nil {
			return nil, err
		}

		if DetectAESInECBMode(decoded, keySize) {
			return decoded, nil

		}

	}

	return nil, nil
}
