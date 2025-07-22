package encoder

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log"
)

func unpad(src []byte) []byte {
	length := len(src)
	if length == 0 {
		return nil
	}
	padding := int(src[length-1])
	if padding > length {
		return nil
	}
	return src[:length-padding]
}

func ivextract(ciphertext []byte) ([]byte, []byte) {

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	return iv, ciphertext
}

func DecoderAES(key []byte, encodedtext []byte) string {
	ciphertext, err := hex.DecodeString(string(encodedtext))
	if err != nil {
		log.Fatal(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	if len(ciphertext) < aes.BlockSize {
		log.Fatal("ciphertext too short")
	}

	iv, ciphertext := ivextract(ciphertext)


	mode := cipher.NewCBCDecrypter(block, iv)

	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)


	plaintext = unpad(plaintext)

	return string(plaintext)
}
