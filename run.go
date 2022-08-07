package main

import (
	"fmt"
	"github.com/gvvsnrnaveen/golang_data_security/base64"
)

func base64_encode_decode(){
	message := "Hai naveen how are you"
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

}

func main(){
	fmt.Println("Main Program");
	
	// base64 sample function
	base64_encode_decode()
}
