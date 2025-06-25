package signature

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

var ErrInvalidSignature = fmt.Errorf("invalid signature")

func VerifySignature(timestamp string, nonce string, body string, secretKey string) string {
	payload := fmt.Sprintf("%s\n%s\n%s\n", timestamp, nonce, body)
	mac := hmac.New(sha512.New, []byte(secretKey))
	mac.Write([]byte(payload))
	signature := mac.Sum(nil)
	return hex.EncodeToString(signature)
}

func Verify(singingData, signature, key string) error {
	sig, err := hex.DecodeString(signature)
	if err != nil {
		return ErrInvalidSignature
	}

	hasher := hmac.New(sha512.New, []byte(key))
	hasher.Write([]byte(singingData))
	if !hmac.Equal(sig, hasher.Sum(nil)) {
		return ErrInvalidSignature
	}

	return nil
}

func Sign(singingData, key string) string {
	hasher := hmac.New(sha512.New, []byte(key))
	hasher.Write([]byte(singingData))
	return hex.EncodeToString(hasher.Sum(nil))
}
