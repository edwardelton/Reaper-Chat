package main

import (
	"TalkerClient/global"
	"TalkerClient/utils"
	"crypto/aes"
	"crypto/cipher"
)

func EncryptText(text []byte) []byte {
	c, err := aes.NewCipher(global.ENCRYPTION_KEY)
	utils.HandleError(err, "Error on the encryption.\nWe will close the client.")

	gcm, err := cipher.NewGCM(c)
	utils.HandleError(err, "Error on the encryption.\nWe will close the client.")

	nonce := make([]byte, gcm.NonceSize())
	utils.HandleError(err, "Error on the encryption.\nWe will close the client.")

	return gcm.Seal(nonce, nonce, text, nil)
}

func DecryptText(text []byte) string {
	c, err := aes.NewCipher(global.ENCRYPTION_KEY)
	utils.HandleError(err, "Error on the encryption.\nWe will close the client.")

	gcm, err := cipher.NewGCM(c)
	utils.HandleError(err, "Error on the encryption.\nWe will close the client.")

	nonceSize := gcm.NonceSize()
	nonce, text := text[:nonceSize], text[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, text, nil)
	utils.HandleError(err, "Error on the encryption.\nWe will close the client.")

	return string(plaintext)
}
