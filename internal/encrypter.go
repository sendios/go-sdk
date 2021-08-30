package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"strings"
)

var key = []byte("444500b1bd43b59f")

type Encrypt struct {
	block cipher.Block
}

func MakeEncrypt() (Encrypt, error) {
	var block cipher.Block
	var err error

	if block, err = aes.NewCipher(key); err != nil {
		return Encrypt{}, err
	}

	return Encrypt{block: block}, nil
}

func (encrypt Encrypt) Decrypt(encryptedData string) (string, error) {
	var err error

	cipherText, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	if len(cipherText) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	if len(cipherText)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}

	cbc := cipher.NewCBCDecrypter(encrypt.block, iv)

	cbc.CryptBlocks(cipherText, cipherText)

	return strings.TrimSpace(string(cipherText)), nil
}

func (encrypt Encrypt) EncryptData(dataToEncrypt []byte) (string, error) {
	encryptedString := base64.StdEncoding.EncodeToString(dataToEncrypt)

	return strings.TrimSpace(encryptedString), nil
}
