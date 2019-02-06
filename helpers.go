package jwt

import (
	"crypto/rsa"
	"io/ioutil"
)

func loadRSAPrivateKeyFromDisk(location string) *rsa.PrivateKey {
	keyData, e := ioutil.ReadFile(location)
	if e != nil {
		panic(e.Error())
	}
	key, e := ParseRSAPrivateKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key
}

func loadRSAPublicKeyFromDisk(location string) *rsa.PublicKey {
	keyData, e := ioutil.ReadFile(location)
	if e != nil {
		panic(e.Error())
	}
	key, e := ParseRSAPublicKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key
}

func makeSampleToken(c Claims, key interface{}) string {
	token := NewWithClaims(SigningMethodRS256, c)
	s, e := token.SignedString(key)

	if e != nil {
		panic(e.Error())
	}

	return s
}
