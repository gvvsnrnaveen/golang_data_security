module github.com/gvvsnrnaveen/golang_data_security

go 1.18

replace github.com/gvvsnrnaveen/golang_data_security/base64 => ./base64

require (
	github.com/gvvsnrnaveen/golang_data_security/aes v0.0.0-00010101000000-000000000000
	github.com/gvvsnrnaveen/golang_data_security/base64 v0.0.0-00010101000000-000000000000
)

require github.com/mergermarket/go-pkcs7 v0.0.0-20170926155232-153b18ea13c9 // indirect

replace github.com/gvvsnrnaveen/golang_data_security/aes => ./aes
