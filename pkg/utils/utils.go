package utils

import (
	"crypto/rand"
	"fmt"
)

const CURRENCY = "€"

//GenerateUUID function: this function does not RFC4122 compliant
func GenerateUUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

func FormatPrice(price float32) string {
	return fmt.Sprintf("%0.02f%s", price, CURRENCY)
}
