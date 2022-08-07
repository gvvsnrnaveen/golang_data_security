package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"github.com/mergermarket/go-pkcs7"
)

func AesEncryption(data []byte, key []byte, iv []byte) ([]byte, error){
	c,err := aes.NewCipher(key)
	if err != nil{
		return nil, err
	}
	data, err = pkcs7.Pad(data, c.BlockSize())
	if err != nil {
		return nil, err
	}
	cbc := cipher.NewCBCEncrypter(c, iv)

	encrypt_data := make([]byte, aes.BlockSize + len(data))
	cbc.CryptBlocks(encrypt_data, data)
	return encrypt_data, nil
}

func AesDecryption(data []byte, key []byte, iv []byte) ([]byte, error){
	c,err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	cbc := cipher.NewCBCDecrypter(c, iv)
	decrypt_data := make([]byte, aes.BlockSize + len(data))
	cbc.CryptBlocks(decrypt_data, data)

	plaintext,err := pkcs7.Unpad(decrypt_data, aes.BlockSize)
	return plaintext, err
}
