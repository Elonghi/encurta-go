package helpers

import (
	"encoding/base64"

	"github.com/skip2/go-qrcode"
)

func GenerateQrCode(url string) (string, error) {
	pngData, err := qrcode.Encode("http://localhost:8080/"+url, qrcode.Medium, 256)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(pngData), nil
}
