package basics

import (
	"encoding/base64"
	"errors"
	"os"
)

func HammingDistance(str1, str2 string) (int, error) {
	if len(str1) != len(str2) {
		return 0, errors.New("strings have different lengths")
	}

	distance := 0
	for i := range str1 {
		b1 := []byte(str1[i : i+1])
		b2 := []byte(str2[i : i+1])
		// The ^ operators sets to 1 only the bits that are different
		val := b1[0] ^ b2[0]

		// We then count the bit set to 1
		for v := val; v > 0; v >>= 1 {
			distance += int(v & 1)
		}
	}

	return distance, nil
}

func FindKeySize(ciphertext []byte) (int, error) {
	minKeySize := 2
	maxKeySize := 40

	bestKeySize := 0
	bestDistance := float64(len(ciphertext))

	for keySize := minKeySize; keySize <= maxKeySize; keySize++ {
		doubleKeySize := keySize * 2
		blocks := len(ciphertext)/doubleKeySize - 1

		if blocks <= 2 {
			// Not enough blocks to calculate a meaningful distance
			continue
		}

		distance := 0

		for block := 0; block < blocks; block++ {
			block1 := ciphertext[block*doubleKeySize : block*doubleKeySize+keySize]
			block2 := ciphertext[block*doubleKeySize+keySize : block*doubleKeySize+2*keySize]

			d, err := HammingDistance(string(block1), string(block2))
			if err != nil {
				return 0, err
			}
			distance += d
		}

		normalizedDistance := (float64(distance) / float64(keySize)) / float64(blocks)

		if normalizedDistance < bestDistance {
			bestDistance = normalizedDistance
			bestKeySize = keySize
		}
	}

	return bestKeySize, nil
}

func decodeBase64File(filename string) ([]byte, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	decoded, err := base64.StdEncoding.DecodeString(string(content))
	if err != nil {
		return nil, err
	}

	return decoded, nil
}

func FindKey(keySize int, ciphertext []byte) []byte {
	// Now that you probably know the KEYSIZE: break the ciphertext into blocks of KEYSIZE length.
	blocks := make([][]byte, keySize)
	for i := 0; i < len(ciphertext); i += keySize {
		block := make([]byte, 0)
		for j := i; j < len(ciphertext); j++ {
			block = append(block, ciphertext[j])
		}
		blocks = append(blocks, block)
	}

	// Now transpose the blocks: make a block that is the first byte of every block, and a
	// block that is the second byte of every block, and so on.
	transposedBlocks := make([][]byte, keySize)
	for i := 0; i < keySize; i++ {
		block := make([]byte, len(blocks))
		for j := 0; j < len(blocks); j++ {
			if i < len(blocks[j]) {
				block[j] = blocks[j][i]
			}
		}
		transposedBlocks[i] = block
	}

	// Solve each block as if it was single-character XOR. You already have code to do this.
	// For each block, the single-byte XOR key that produces the best looking histogram is the
	// repeating-key XOR key byte for that block. Put them together and you have the key.
	theKey := make([]byte, 0)
	for _, block := range transposedBlocks {
		bestKeyByte := byte(0)
		bestScore := 0.0
		for k := range 256 {
			key := byte(k)
			decrypted := singleByteXOR(block, key)

			score := calculateScore(decrypted)

			if score > bestScore {
				bestScore = score
				bestKeyByte = key
			}
		}
		theKey = append(theKey, bestKeyByte)
	}

	return theKey
}

func BreakRepeatingKeyXOR(filename string) ([]byte, error) {
	ciphertext, err := decodeBase64File(filename)
	if err != nil {
		return nil, err
	}

	keySize, err := FindKeySize(ciphertext)
	if err != nil {
		return nil, err
	}

	key := FindKey(keySize, ciphertext)
	decrypted := RepeatingKeyXOR(ciphertext, key)

	return decrypted, nil
}
