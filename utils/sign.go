package utils

import (
	"crypto/ecdsa"
	cryptoBee "github.com/ethersphere/bee/pkg/crypto"
)

func Sign(messageId []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	signer := cryptoBee.NewDefaultSigner(privateKey)
	signature, err := signer.Sign(messageId)
	if err != nil {
		return nil, err
	}

	return signature, nil
	//return hexutil.Encode(signature), nil
}
