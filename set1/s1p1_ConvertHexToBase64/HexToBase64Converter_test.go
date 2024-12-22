package hextobase64converter

import (
	"testing"

	"github.com/phuongdoan13/Cryptopals/set1/utils"
)

func TestConvertHexStrToBase64Str(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	result, err := utils.HexStrToBase64Str(input)
	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf("Wrong hex to 64 conversion")
	}
}