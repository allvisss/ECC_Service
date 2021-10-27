package keygen

import (
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	"github.com/allvisss/ECC_service/common/model"
	"github.com/allvisss/ECC_service/common/response"

	ecies "github.com/ecies/go"
	"github.com/fomichev/secp256k1"
)

type (
	IKeyGenerate interface {
		GenerateKey(curve model.Curve)
	}
	KeyPair struct {
		PublicKey  string
		PrivateKey string
	}
	secp112r1Curve struct {
		*elliptic.CurveParams
	}
)

func GenerateKey() (int, interface{}) {
	priv, err := GeneratePriv()
	if err != nil {
		return response.ServiceUnavailableMsg("Cannot generate key.")
	}
	err = WriteKeyToFile(priv)
	if err != nil {
		return response.ServiceUnavailable()
	}
	return response.NewOKResponse(priv)
}

func GeneratePriv() (*model.PrivateKey, error) {
	curve := secp256k1.SECP256K1()

	p, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("cannot generate key pair: %w", err)
	}

	return &model.PrivateKey{
		PublicKey: &model.PublicKey{
			Curve: curve,
			X:     x,
			Y:     y,
		},
		D: new(big.Int).SetBytes(p),
	}, nil
}

// func setCurveSecp112r1() elliptic.Curve {
// 	curveParams := &elliptic.CurveParams{Name: "secp112r1"}
// 	curveParams.P, _ = new(big.Int).SetString("DB7C2ABF62E35E668076BEAD208B", 16)
// 	curveParams.B, _ = new(big.Int).SetString("659EF8BA043916EEDE8911702B22", 16)
// 	curveParams.Gx, _ = new(big.Int).SetString("09487239995A5EE76B55F9C2F098", 16)
// 	curveParams.Gy, _ = new(big.Int).SetString("A89CE5AF8724C0A23E0E0FF77500", 16)
// 	curveParams.N, _ = new(big.Int).SetString("DB7C2ABF62E35E7628DFAC6561C5", 16)
// 	curveParams.BitSize = 112
// 	curve := secp112r1Curve{curveParams}
// 	return curve
// }

func WriteKeyToFile(priv *model.PrivateKey) error {
	f, err := os.Create("privatekey.txt")
	if err != nil {
		return fmt.Errorf("create file failed: %w", err)
	}

	defer f.Close()

	data, err := json.Marshal(priv)
	if err != nil {
		return fmt.Errorf("write failed: %w", err)
	}
	_, err = f.Write(data)
	if err != nil {
		return fmt.Errorf("write failed: %w", err)
	}
	return nil
}

func WriteKeyToFile2(priv *ecies.PrivateKey) error {
	f, err := os.Create("privatekey2.txt")
	if err != nil {
		return fmt.Errorf("create file failed: %w", err)
	}

	defer f.Close()

	data, err := json.Marshal(priv)
	if err != nil {
		return fmt.Errorf("write failed: %w", err)
	}
	_, err = f.Write(data)
	if err != nil {
		return fmt.Errorf("write failed: %w", err)
	}
	return nil
}
