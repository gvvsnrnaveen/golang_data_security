package base64

import (
	"encoding/base64"
)

func Base64Encoding(data []byte) (string){
	encoded := base64.StdEncoding.EncodeToString(data)
	return encoded
}

func Base64Decoding(data string) ([]byte, error){
	decoded,err := base64.StdEncoding.DecodeString(data)
	return decoded,err
}
