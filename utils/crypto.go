package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
    "encoding/base64"
	"io/ioutil"
	"os"
)

func GenerateKeys() error {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return err
	}

	privKeyBytes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return err
	}
	privKeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: privKeyBytes,
	})

	pubKeyBytes, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
        return err
	}
	pubKeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyBytes,
	})

	privKeyFile, err := os.Create("private.pem")
	if err != nil {
        return err
	}
	defer privKeyFile.Close()
	_, err = privKeyFile.Write(privKeyPem)
	if err != nil {
        return err
	}

	pubKeyFile, err := os.Create("public.pem")
	if err != nil {
        return err
	}
	defer pubKeyFile.Close()
	_, err = pubKeyFile.Write(pubKeyPem)
	if err != nil {
        return err
	}
    return nil
}

func SignMessage(message []byte) (string, error) {
    pk, err := loadPrivateKey()
    if err != nil {
        return "", err
    }

    r, s, err := ecdsa.Sign(rand.Reader, pk, message) 
    if err != nil {
        return "", nil
    }

    signature := append(r.Bytes(), s.Bytes()...)
    res := base64.StdEncoding.EncodeToString(signature)
    return res, nil
}

func loadPrivateKey() (*ecdsa.PrivateKey, error) {
    keyData, err := ioutil.ReadFile("private.pem")
    if err != nil {
        return nil, err
    }

    block, _ := pem.Decode(keyData)
    if block == nil {
        return nil, errors.New("Could not decode pem file")
    }

    privKey, err := x509.ParseECPrivateKey(block.Bytes)
    if err != nil {
        return nil, err
    }

    return privKey, nil
}
