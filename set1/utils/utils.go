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
// Example:
// 	hexStr = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
// 	expected = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
func HexStrToBase64Str(hexStr string) (string, error) {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}