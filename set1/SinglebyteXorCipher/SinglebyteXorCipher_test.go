package singlebyteXorCipher

import (
	"encoding/hex"
	"testing"
)

func TestSinglebyteXorCipherWithEnglishOnly(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	expected := "Cooking MC's like a pound of bacon"
	result, err := SinglebyteXorCipher(input)
	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf("\nTest fail. Decipher is not correct.\nExpect: \t '%s' \nbut get:\t '%s'", expected, result)
	}
}

func TestSinglebyteXorCipherWithMultiByteCharacters(t *testing.T) {
	expected := "Hello, XOR Cipher!"
	input := createInputStrFromExpected(expected)

	result, err := SinglebyteXorCipher(input)
	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf("\nTest fail. Decipher is not correct.\nExpect: \t '%s' \nbut get:\t '%s'", expected, result)
	}
}

func createInputStrFromExpected(expected string) string {
	cipherTextBytes := xorCipher([]byte(expected), byte('2'))
	return hex.EncodeToString(cipherTextBytes)
}

func xorCipher(input []byte, key byte) []byte {
	result := make([]byte, len(input))
	for i := range input {
		result[i] = input[i] ^ key // XOR each byte with the key
	}
	return result
}
