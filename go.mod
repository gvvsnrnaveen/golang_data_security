module github.com/gvvsnrnaveen/golang_data_security

go 1.18

replace github.com/gvvsnrnaveen/golang_data_security/base64 => ./base64

require (
	github.com/gvvsnrnaveen/golang_data_security/aes v0.0.0-00010101000000-000000000000
	github.com/gvvsnrnaveen/golang_data_security/base64 v0.0.0-00010101000000-000000000000
	github.com/gvvsnrnaveen/golang_data_security/sha v0.0.0-00010101000000-000000000000
)

require (
	github.com/mergermarket/go-pkcs7 v0.0.0-20170926155232-153b18ea13c9 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
)

replace github.com/gvvsnrnaveen/golang_data_security/aes => ./aes

replace github.com/gvvsnrnaveen/golang_data_security/sha => ./sha
