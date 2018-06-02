package core

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

func EncryptRSA(msg []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	hash := sha256.New()

	encrypted, err := rsa.EncryptOAEP(hash, rand.Reader, &privateKey.PublicKey, msg, nil)
	if err != nil {
		return nil, err
	}
	return encrypted, nil
}

func DencryptRSA(cipher []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	hash := sha256.New()

	decrypted, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, cipher, nil)
	if err != nil {
		return nil, err
	}
	return decrypted, nil
}

func ReadPrivateKey(path string) (*rsa.PrivateKey, string, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, "", err
	}

	pdat := string(dat)
	block, _ := pem.Decode([]byte(pdat))
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	return key, SHA256(dat), nil
}
