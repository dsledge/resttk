package resttk

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
)

var (
	Secret string
)

func secret() string {
	if len(Secret) <= 1 {
		id := make([]byte, 32)
		if _, err := rand.Read(id); err != nil {
			Secret = ""
		}
		Secret = fmt.Sprintf("%x", id)
	}
	return Secret
}

func GenerateAuthToken(message string) string {
	mac := hmac.New(sha256.New, []byte(secret()))
	mac.Write([]byte(message + time.Now().String()))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
