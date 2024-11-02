package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d" 
	output, err := ConvertHexToBase64(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output) 
}

func ConvertHexToBase64(hexStr string) (string, error){
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}