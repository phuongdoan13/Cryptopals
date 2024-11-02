package fixedxor

import "testing"

func TestXorHexToGetBase64String(t *testing.T) {
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"
	result, err := XorHexStr(input1, input2)
	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatal("Wrong xor operation")
	}
}
