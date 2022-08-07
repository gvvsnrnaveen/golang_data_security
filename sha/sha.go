package sha

import (
	"crypto/sha256"
	"crypto/sha512"
	"golang.org/x/crypto/sha3"
)

func PerformSha(message []byte) ([]string, []string){
	result := []string {}
	engine := []string {}

	// sha256
	sha := sha256.New()
	sha.Write(message)
	result = append(result, string(sha.Sum(nil)))
	engine = append(engine, "SHA256")

	// sha384
	sha = sha512.New384()
	sha.Write(message)
	result = append(result, string(sha.Sum(nil)))
	engine = append(engine, "SHA384")

	// sha512
	sha = sha512.New()
	sha.Write(message)
	result = append(result, string(sha.Sum(nil)))
	engine = append(engine, "SHA512")

	// sha3-256
	sha = sha3.New256()
	sha.Write(message)
	result = append(result, string(sha.Sum(nil)))
	engine = append(engine, "SHA3-256")

	// sha3-384
	sha = sha3.New384()
	sha.Write(message)
	result = append(result, string(sha.Sum(nil)))
	engine = append(engine, "SHA3-384")

	// sha3-512
	sha = sha3.New512()
	sha.Write(message)
	result = append(result, string(sha.Sum(nil)))
	engine = append(engine, "SHA3-512")

	return engine,result
}
