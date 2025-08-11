package utils

import (
	"encoding/base64"
	"encoding/hex"
)

// Convert hex string to base64 string
//
// Parameters:
// - hexStr (string): a string in hex representation
// 
// Returns:
// - (string): the resulted string in base64 represetnation
//
// Example:
// - hexStr = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
// - expected = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
func HexStrToBase64Str(hexStr string) (string, error) {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

// Implements a repeating key XOR encryption algorithm. It takes a plaintext input string and a key string, applies the XOR operation to each character of the input string using the corresponding character from the repeating key, and then returns the result as a hexadecimal-encoded string.
//
// Parameters:
// - input (string): a human-readable string
// - key (string): a key used for encoding
//
// Returns:
// - (string): resulted hex
//
// Examples:
// - input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
// - key := "ICE"
// - expected := `0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`
func RepeatingKeyXor(input string, key string) string {
	inputBytes := []byte(input)
	keyBytes := []byte(key)
	
	ansBytes := make([]byte, len(inputBytes))

	for i := 0; i < len(input); i++ {
		ansBytes[i] = inputBytes[i] ^ keyBytes[i % len(key)]
	}

	return hex.EncodeToString(ansBytes)
}