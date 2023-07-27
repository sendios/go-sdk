package internal

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func Base64Encoder(toEncode string) string {
	return base64.StdEncoding.EncodeToString([]byte(toEncode))
}

func Sha1Encoder(toEncode string) string {
	s := sha1.New()
	s.Write([]byte(toEncode))
	key := s.Sum(nil)

	return fmt.Sprintf("%x", key)
}
