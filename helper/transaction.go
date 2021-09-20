package helper

import (
	"bitacoin_client/schema"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"fmt"
	"time"
)

func CreateTxnReq(prKey, pubKey, toPubKey []byte, amount int) (*schema.TransactionRequest, error) {
	pr, err := x509.ParsePKCS1PrivateKey(prKey)
	if err != nil {
		return nil, err
	}

	tnxTime := time.Now().UnixNano()
	hasher := sha256.New()
	_, err = fmt.Fprint(hasher, tnxTime, pubKey, toPubKey, amount)
	if err != nil {
		return nil, err
	}
	hashed := hasher.Sum(nil)

	sig, err := rsa.SignPKCS1v15(rand.Reader, pr, crypto.SHA256, hashed)
	if err != nil {
		return nil, err
	}

	t := schema.TransactionRequest{
		Time:       tnxTime,
		FromPubKey: pubKey,
		ToPubKey:   toPubKey,
		Signature:  sig,
		Amount:     amount,
	}

	return &t, nil
}
