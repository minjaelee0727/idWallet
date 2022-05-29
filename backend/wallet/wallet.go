package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"math/big"

	"github.com/minjaelee0727/idWallet/backend/utils"
)

// IdWallt only contains keys
type IdWallet struct {
	PrivateKey string
	PublicKey  string
}

var w *IdWallet

func createPrivKey() *ecdsa.PrivateKey {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)
	return privKey
}

func encode(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return string(pemEncoded), string(pemEncodedPub)
}

func decode(pemEncoded string, pemEncodedPub string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	block, _ := pem.Decode([]byte(pemEncoded))
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)

	blockPub, _ := pem.Decode([]byte(pemEncodedPub))
	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericPublicKey.(*ecdsa.PublicKey)

	return privateKey, publicKey
}

func restoreBigInts(payload string) (*big.Int, *big.Int, error) {
	bytes, err := hex.DecodeString(payload)
	if err != nil {
		return nil, nil, err
	}
	firstHalfBytes := bytes[:len(bytes)/2]
	sencodHalfBytes := bytes[len(bytes)/2:]
	bigA, bigB := big.Int{}, big.Int{}
	bigA.SetBytes(firstHalfBytes)
	bigB.SetBytes(sencodHalfBytes)
	return &bigA, &bigB, nil
}

func MakeSignature(hashedCredential string, w *IdWallet) string {
	payloadByte, err := hex.DecodeString(hashedCredential)
	utils.HandleErr(err)
	privKey, _ := decode(w.PrivateKey, w.PublicKey)
	r, s, err := ecdsa.Sign(rand.Reader, privKey, payloadByte)
	utils.HandleErr(err)
	signBytes := append(r.Bytes(), s.Bytes()...)
	return fmt.Sprintf("%x", signBytes)
}

func VerifySignature(signature, publicKeyStr, hashedCredential string) bool {
	r, s, err := restoreBigInts(signature)
	utils.HandleErr(err)
	x, y, err := restoreBigInts(publicKeyStr)
	utils.HandleErr(err)
	publicKey := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}
	credentialBytes, err := hex.DecodeString(hashedCredential)
	utils.HandleErr(err)
	result := ecdsa.Verify(&publicKey, credentialBytes, r, s)
	return result
}

func CreateWallet() *IdWallet {
	privKey := createPrivKey()
	privKeyString, pubKeyString := encode(privKey, &privKey.PublicKey)
	w = &IdWallet{
		PrivateKey: privKeyString,
		PublicKey:  pubKeyString,
	}
	return w
}
