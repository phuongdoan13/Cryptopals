package implementRepeatingKeyXor

import "encoding/hex"

func RepeatingKeyXor(input string, key string) string {
	inputBytes := []byte(input)
	keyBytes := []byte(key)
	
	ansBytes := make([]byte, len(inputBytes))

	for i := 0; i < len(input); i++ {
		ansBytes[i] = inputBytes[i] ^ keyBytes[i % len(key)]
	}

	return hex.EncodeToString(ansBytes)
}