package main

import (
	"fmt"
	"github.com/gvvsnrnaveen/golang_data_security/base64"
	"github.com/gvvsnrnaveen/golang_data_security/aes"
)

func base64_encode_decode(message string){
	fmt.Println("Base64 encode/decode")
	fmt.Println("Message: ", message);

	// base64 encoding
	encoded := base64.Base64Encoding([]byte(message))
	fmt.Println("Base64 encoded string: ", encoded)

	// base64 decoding
	decoded,err := base64.Base64Decoding(encoded)
	if err != nil {
		fmt.Println("Failed: ", err)
		return
	}
	fmt.Println("Decoded string: ", string(decoded))
	fmt.Println("===========================\n")

}

func aes_encrypt_decrypt(message string){
	key:= "1234567812345678"
	iv:= "1234567812345678"
	fmt.Println("AES Encryption/Decryption")
	cipher,err := aes.AesEncryption([]byte(message), []byte(key), []byte(iv))
	if err != nil{
		fmt.Println("Failed to perform encryption: ", err)
		return
	}
	encoded := base64.Base64Encoding(cipher)
	fmt.Println("Cipher: ", encoded)

	plaintext,err := aes.AesDecryption(cipher, []byte(key), []byte(iv))
	if err != nil {
		fmt.Println("Failed to perform decryption: ", err)
		return
	}
	fmt.Println("Decrypted: ", string(plaintext[:len(message)]))
	fmt.Println("===========================\n");
}

func main(){
	fmt.Println("Main Program");
	message := "Hai naveen how are you"
	
	// base64 sample function
	base64_encode_decode(message)

	// aes encryption decryption sample
	aes_encrypt_decrypt(message)
}
