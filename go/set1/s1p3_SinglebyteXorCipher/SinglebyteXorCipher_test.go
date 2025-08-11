package singlebyteXorCipher

import (
	"encoding/hex"
	"fmt"
	"testing"
)

// Question:
// Single-byte XOR cipher
// The hex encoded string:

// 1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
// ... has been XOR'd against a single character. Find the key, decrypt the message.

// You can do this by hand. But don't: write code to do it for you.

// How? Devise some method for "scoring" a piece of English plaintext. Character frequency is a good metric. Evaluate each output and choose the one with the best score.

// Achievement Unlocked
// You now have our permission to make "ETAOIN SHRDLU" jokes on Twitter.

// Reflection:
// The test does not support text with multi-byte characters and punctuations

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

type TestCase struct {
	input string
	expected string
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []TestCase {
		{createInputStrFromExpected("Hello world"), "Hello world"},
		{createInputStrFromExpected("Welcome to cryptopal"), "Welcome to cryptopal"},
		{createInputStrFromExpected("This is a test that you need! @harry"), "This is a test that you need! @harry"},
		{createInputStrFromExpected("I'm so happy that I have a chance to read on O'reilly"), "I'm so happy that I have a chance to read on O'reilly"},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		testname := fmt.Sprintf("%s,%s", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			ans, err := SinglebyteXorCipher(tt.input)
			if err != nil {
				t.Fatal(err)
			}

			if ans != tt.expected {
				t.Errorf("Got:	'%s'\nExpected: '%s'", 
				ans, tt.expected)
			}
		})
	}
}

func TestSinglebyteXorCipherWithCustomTestcases(t *testing.T) {
	expected := "Hello my name is Harry."
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
	cipherTextBytes := xorCipher([]byte(expected), byte('a'))
	return hex.EncodeToString(cipherTextBytes)
}

func xorCipher(input []byte, key byte) []byte {
	result := make([]byte, len(input))
	for i := range input {
		result[i] = input[i] ^ key // XOR each byte with the key
	}
	return result
}
