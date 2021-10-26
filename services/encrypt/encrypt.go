package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"

	"github.com/allvisss/ECC_service/common/model"
	"github.com/allvisss/ECC_service/common/response"
	"github.com/allvisss/ECC_service/services/keygen"
)

func EncryptService(msg string, priv model.PrivateKey) (int, interface{}) {
	ciphertext, err := Encrypt(priv.PublicKey, []byte(msg))
	if err != nil {
		return response.ServiceUnavailableMsg("Can not encrypt")
	}

	return response.NewOKResponse(ciphertext)
}

func Encrypt(pubkey *model.PublicKey, msg []byte) ([]byte, error) {
	var ct bytes.Buffer

	// Generate ephemeral key
	ek, err := keygen.GeneratePriv()
	if err != nil {
		return nil, err
	}

	ct.Write(ek.PublicKey.Bytes(false))

	// Derive shared secret
	ss, err := ek.Encapsulate(pubkey)
	if err != nil {
		return nil, err
	}

	// AES encryption
	block, err := aes.NewCipher(ss)
	if err != nil {
		return nil, fmt.Errorf("cannot create new aes block: %w", err)
	}

	nonce := make([]byte, 16)
	if _, err := rand.Read(nonce); err != nil {
		return nil, fmt.Errorf("cannot read random bytes for nonce: %w", err)
	}

	ct.Write(nonce)

	aesgcm, err := cipher.NewGCMWithNonceSize(block, 16)
	if err != nil {
		return nil, fmt.Errorf("cannot create aes gcm: %w", err)
	}

	ciphertext := aesgcm.Seal(nil, nonce, msg, nil)

	tag := ciphertext[len(ciphertext)-aesgcm.NonceSize():]
	ct.Write(tag)
	ciphertext = ciphertext[:len(ciphertext)-len(tag)]
	ct.Write(ciphertext)

	return ct.Bytes(), nil
}
