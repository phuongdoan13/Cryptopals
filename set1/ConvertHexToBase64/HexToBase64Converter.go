package hextobase64converter

import (
	"encoding/base64"
	"encoding/hex"
)

func ConvertHexStrToBase64Str(hexStr string) (string, error) {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}
