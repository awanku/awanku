package utils

import (
	"crypto/hmac"
	"crypto/sha256"
)

func ValidateHmac(key, plain, hashed []byte) (bool, error) {
	mac := hmac.New(sha256.New, key)
	_, err := mac.Write(plain)
	if err != nil {
		return false, err
	}
	computed := mac.Sum(nil)
	return hmac.Equal(computed, hashed), nil
}
