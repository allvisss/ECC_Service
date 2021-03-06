package model

import (
	"bytes"
	"crypto/elliptic"
	"fmt"
	"math/big"
)

type PublicKey struct {
	elliptic.Curve
	X, Y *big.Int
}

// Bytes returns public key raw bytes;
// Could be optionally compressed by dropping Y part
func (k *PublicKey) Bytes(compressed bool) []byte {
	x := k.X.Bytes()
	if len(x) < 32 {
		for i := 0; i < 32-len(x); i++ {
			x = append([]byte{0}, x...)
		}
	}

	if compressed {
		// If odd
		if k.Y.Bit(0) != 0 {
			return bytes.Join([][]byte{{0x03}, x}, nil)
		}

		// If even
		return bytes.Join([][]byte{{0x02}, x}, nil)
	}

	y := k.Y.Bytes()
	if len(y) < 32 {
		for i := 0; i < 32-len(y); i++ {
			y = append([]byte{0}, y...)
		}
	}

	return bytes.Join([][]byte{{0x04}, x, y}, nil)
}

// Decapsulate decapsulates key by using Key Encapsulation Mechanism and returns symmetric key;
// can be safely used as encryption key
func (k *PublicKey) Decapsulate(priv *PrivateKey) ([]byte, error) {
	if priv == nil {
		return nil, fmt.Errorf("public key is empty")
	}

	var secret bytes.Buffer
	secret.Write(k.Bytes(false))

	sx, sy := priv.Curve.ScalarMult(k.X, k.Y, priv.D.Bytes())
	secret.Write([]byte{0x04})

	// Sometimes shared secret coordinates are less than 32 bytes; Big Endian
	l := len(priv.Curve.Params().P.Bytes())
	secret.Write(zeroPad(sx.Bytes(), l))
	secret.Write(zeroPad(sy.Bytes(), l))

	return kdf(secret.Bytes())
}
