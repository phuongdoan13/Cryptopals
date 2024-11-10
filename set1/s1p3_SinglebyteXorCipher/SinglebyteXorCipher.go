package singlebyteXorCipher

import (
	"encoding/hex"
	"sort"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

func SinglebyteXorCipher(input string) (string, error) {
	cipherText, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}
	return string(getHighestPossiblePlaintext(cipherText)), nil	
}

type BFXOR struct {
	score float64
	prop float64
	key byte
	res []byte
}

func getHighestPossiblePlaintext(input []byte) []byte {
	ref := make([]BFXOR, 256)

	for i := 0; i < 256; i++ {
		plaintextBytes := getPlaintextBytes(input, byte(i))
		score, prop := scoreEnglishCharacter(plaintextBytes)

		ref[i] = BFXOR{
			score: score,
			prop: prop,
			key: byte(i),
			res: plaintextBytes,
		}
	}

	// the lower the chi-square score, the closer the observed distribution to the expected distribution
	sort.Slice(ref, func(i, j int) bool {return ref[i].score < ref[j].score})

	return ref[0].res
}

func getPlaintextBytes(cipherText []byte, key byte) []byte {
	result := make([]byte, len(cipherText))
	for i := range cipherText {
		result[i] = cipherText[i] ^ key
	}
	return result
}

func scoreEnglishCharacter(input []byte) (float64, float64) {
	// ref: https://github.com/mattnotmitt/CryptoPals-Go/blob/master/set1/chal3.go
	// the idea is to use chi-squared metric to tell if the the decipher result has the distribution close that of normal English
	engFreq := []float64{0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.755, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 15.843, 0.004, 0.375, 0.002, 0.008, 0.019, 0.008, 0.134, 0.137, 0.137, 0.001, 0.001, 0.972, 0.19, 0.857, 0.017, 0.334, 0.421, 0.246, 0.108, 0.104, 0.112, 0.103, 0.1, 0.127, 0.237, 0.04, 0.027, 0.004, 0.003, 0.004, 0.002, 0.0001, 0.338, 0.218, 0.326, 0.163, 0.121, 0.149, 0.133, 0.192, 0.232, 0.107, 0.082, 0.148, 0.248, 0.134, 0.103, 0.195, 0.012, 0.162, 0.368, 0.366, 0.077, 0.061, 0.127, 0.009, 0.03, 0.015, 0.004, 0.0001, 0.004, 0.0001, 0.003, 0.0001, 6.614, 1.039, 2.327, 2.934, 9.162, 1.606, 1.415, 3.503, 5.718, 0.081, 0.461, 3.153, 1.793, 5.723, 5.565, 1.415, 0.066, 5.036, 4.79, 6.284, 1.992, 0.759, 1.176, 0.139, 1.162, 0.102, 0.0001, 0.002, 0.0001, 0.0001, 0.0001, 0.06, 0.004, 0.003, 0.002, 0.001, 0.001, 0.001, 0.002, 0.001, 0.001, 0.0001, 0.001, 0.001, 0.003, 0.0001, 0.0001, 0.001, 0.001, 0.001, 0.031, 0.006, 0.001, 0.001, 0.001, 0.002, 0.014, 0.001, 0.001, 0.005, 0.005, 0.001, 0.002, 0.017, 0.007, 0.002, 0.003, 0.004, 0.002, 0.001, 0.002, 0.002, 0.012, 0.001, 0.002, 0.001, 0.004, 0.001, 0.001, 0.003, 0.003, 0.002, 0.005, 0.001, 0.001, 0.003, 0.001, 0.003, 0.001, 0.002, 0.001, 0.004, 0.001, 0.002, 0.001, 0.0001, 0.0001, 0.02, 0.047, 0.009, 0.009, 0.0001, 0.0001, 0.001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.003, 0.001, 0.004, 0.002, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.001, 0.001, 0.001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.005, 0.002, 0.061, 0.001, 0.0001, 0.002, 0.001, 0.001, 0.001, 0.001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001}

	// calc frequency of each ascii character in the text
	counts := make([]int, 256)
	for _, b := range input {
		counts[b]++
	}

	// calc frequency of distribution for each 
	freqDist := make([]float64, 256)
	for i, c := range counts {
		freqDist[i] = (float64(c) / float64(256)) * 100
	}

	// calc stats
	score := stat.ChiSquare(freqDist, engFreq)
	df := float64(len(freqDist) - 1)
	return score, 1 - distuv.ChiSquared{K: df}.CDF(score)
}