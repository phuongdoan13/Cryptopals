package fixedxor

import (
	"encoding/hex"
	"errors"
)

func XorHexStr(input1 string, input2 string) (string, error) {
	bytes1, err := hex.DecodeString(input1)
	if err != nil {
		return "", err
	}

	bytes2, err := hex.DecodeString(input2)
	if err != nil {
		return "", err
	}

	resultBytes, err := xorBytes(bytes1, bytes2)
	if err != nil {
		return "", err
	}

	result := hex.EncodeToString(resultBytes)
	return result, nil
}

func xorBytes(bytes1 []byte, bytes2 []byte) ([]byte, error) {

	if len(bytes1) != len(bytes2) {
		return nil, errors.New("bytes are not the same length")
	}

	result := make([]byte, len(bytes1))

	for i := range bytes1 {
		result[i] = bytes1[i] ^ bytes2[i]
	}

	return result, nil
}
