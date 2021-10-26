package model

import "math/big"

type Curve struct {
	P      big.Int // order of underlying field
	A      big.Int // parameter a in y^2 = x^3 + ax + b
	B      big.Int // parameter b in y^2 = x^3 + ax + b
	Gx, Gy big.Int // generation point (uncompressed form)
	N      big.Int // order of base point
	H      big.Int // cofactor
}
