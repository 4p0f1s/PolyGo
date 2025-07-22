// modules/encoder/encoder.go
package encoder

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"
)

func pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func ivrand() []byte {
	// Crear IV aleatorio
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Fatal(err)
	}
	return iv
}

func EncoderAES(key []byte, plaintext string) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	paddedText := pad([]byte(plaintext), aes.BlockSize)

	iv := ivrand()

	// Usar modo CBC
	mode := cipher.NewCBCEncrypter(block, iv)

	ciphertext := make([]byte, len(paddedText))
	mode.CryptBlocks(ciphertext, paddedText)

	// Prepend IV al resultado
	finalOutput := append(iv, ciphertext...)

	return hex.EncodeToString(finalOutput)
}
