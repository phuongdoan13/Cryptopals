package singlebyteXorCipher

import (
	"encoding/hex"
)

func SinglebyteXorCipher(input string) (string, error) {
	cipherText, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}
	return string(getHighestPossiblePlaintext(cipherText)), nil	
}

func getHighestPossiblePlaintext(input []byte) []byte {
	var highestEnglishCharacterIndex int;
	var mostPossiblePlaintext []byte;

	for key := 0; key < 256; key++ {
		plaintextBytes := getPlaintextBytes(input, byte(key))
		englishCharacterIndex := getEnglishCharacterIndex(plaintextBytes)

		if (englishCharacterIndex > highestEnglishCharacterIndex) {
			highestEnglishCharacterIndex = englishCharacterIndex
			mostPossiblePlaintext = plaintextBytes
		}
	}
	return mostPossiblePlaintext
}

func getPlaintextBytes(cipherText []byte, key byte) []byte {
	result := make([]byte, len(cipherText))
	for i := range cipherText {
		result[i] = getPlaintext(cipherText[i], key)
	}
	return result
}

func getPlaintext(cipherText byte, key byte) byte {
	return cipherText ^ key
}

func getEnglishCharacterIndex(input []byte) int {
	score := 0
	for _, b := range input {
		if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || b == ' ' {
			score++
		}
	}
	return score
}