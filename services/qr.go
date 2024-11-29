package services

import (
	"github.com/skip2/go-qrcode"
)

func CreateQRCode(data string) ([]byte, error) {
	qr, err := qrcode.Encode(data, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return qr, nil
}
