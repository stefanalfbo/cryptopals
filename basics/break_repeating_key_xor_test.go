package basics_test

import (
	"testing"

	"github.com/stefanalfbo/cryptopals/basics"
)

func TestHammingDistance(t *testing.T) {
	str1 := "test"
	str2 := "test"
	expectedDistance := 0
	distance, err := basics.HammingDistance(str1, str2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if distance != expectedDistance {
		t.Errorf("Expected distance: %d, got: %d", expectedDistance, distance)
	}

	str1 = "this is a test"
	str2 = "wokka wokka!!!"
	expectedDistance = 37
	distance, err = basics.HammingDistance(str1, str2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if distance != expectedDistance {
		t.Errorf("Expected distance: %d, got: %d", expectedDistance, distance)
	}
}

func TestFindKeySize(t *testing.T) {
	plaintext := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")

	ciphertext := basics.RepeatingKeyXOR(plaintext, []byte("ICE"))

	keySize, err := basics.FindKeySize(ciphertext)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if keySize != 3 {
		t.Errorf("Expected key size: %d, got: %d", 3, keySize)
	}
}

func TestFindKey(t *testing.T) {
	actualKey := []byte("ICE")
	plaintext := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")

	ciphertext := basics.RepeatingKeyXOR(plaintext, actualKey)

	key := basics.FindKey(3, ciphertext)
	if string(key) != string(actualKey) {
		t.Errorf("Expected key: %s, got: %s", string(actualKey), string(key))
	}
}
