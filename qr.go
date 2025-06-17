package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/skip2/go-qrcode"
)

func createQRCodeAsPNG(body string, recoveryLevel string) ([]byte, error) {
	if body == "" {
		return nil, errors.New("body must not be empty")
	}

	png, err := qrcode.Encode(
		body,
		parseQRCodeParseRecoveryLevel(recoveryLevel),
		-16,
	)

	if err != nil {
		return nil, err
	}

	return png, nil
}

func getQRCodeFilename() string {
	return fmt.Sprintf(
		"qr_code_%s.png",
		time.Now().Format("20060102_150405"),
	)
}

func parseQRCodeParseRecoveryLevel(recoveryLevel string) qrcode.RecoveryLevel {
	if recoveryLevel == "low" {
		return qrcode.Low
	}

	if recoveryLevel == "high" {
		return qrcode.High
	}

	if recoveryLevel == "highest" {
		return qrcode.Highest
	}

	return qrcode.Medium
}
