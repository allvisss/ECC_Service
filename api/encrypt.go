package api

import (
	"crypto/elliptic"
	"math/big"
	"net/http"

	"github.com/allvisss/ECC_service/common/model"
	"github.com/allvisss/ECC_service/common/response"
	service "github.com/allvisss/ECC_service/services/encrypt"
	ecies "github.com/ecies/go"
	"github.com/gin-gonic/gin"
)

type EncryptHandler struct {
}

type CurveWithParams struct {
	*elliptic.CurveParams
}

type Keypair struct {
	Curve struct {
		P       string `json:"P"`
		N       string `json:"N"`
		B       string `json:"B"`
		Gx      string `json:"Gx"`
		Gy      string `json:"Gy"`
		BitSize int    `json:"BitSize"`
		Name    string `json:"Name"`
	}
	X string `json:"X"`
	Y string `json:"Y"`
	D string `json:"D"`
}

func NewEncryptHandler(r *gin.Engine) {
	handler := &EncryptHandler{}
	v1 := r.Group("/v1")
	{
		v1.POST("/encrypt", handler.Encrypt)
		v1.POST("/encrypt2", handler.Encrypt2)
	}
}

//Encrypt
// @Summary Encrypt plaintext into ciphertext (using curve secp112r1 as demo)
// @Tags Encrypt
// @Accept  json
// @Produce  json
// @Param msg query string false "Plaintext"
// @Param priv body Keypair true "Keypair"
// @Router /v1/encrypt [post]
// @Success 201 {object} response.Response
// @Failure 400,404,500 {object} response.ErrorResponse
func (data *EncryptHandler) Encrypt(c *gin.Context) {
	msg := c.Query("msg")
	var input Keypair
	err := c.BindJSON(&input)
	if err != nil {
		code, result := response.BadRequest()
		c.JSON(code, result)
		return
	}

	var publicKey model.PublicKey
	var priv model.PrivateKey
	curveParams := &elliptic.CurveParams{Name: input.Curve.Name}
	curveParams.P, _ = new(big.Int).SetString(input.Curve.P, 10)
	curveParams.B, _ = new(big.Int).SetString(input.Curve.B, 10)
	curveParams.Gx, _ = new(big.Int).SetString(input.Curve.Gx, 10)
	curveParams.Gy, _ = new(big.Int).SetString(input.Curve.Gy, 10)
	curveParams.N, _ = new(big.Int).SetString(input.Curve.N, 10)
	curveParams.BitSize = 112

	publicKey.Curve = CurveWithParams{curveParams}

	publicKey.X, _ = new(big.Int).SetString(input.X, 10)
	publicKey.Y, _ = new(big.Int).SetString(input.Y, 10)

	priv.PublicKey = &publicKey
	priv.D, _ = new(big.Int).SetString(input.D, 10)

	code, result := service.EncryptService(msg, priv)
	c.JSON(code, result)
}

//Encrypt2
// @Summary Encrypt plaintext into ciphertext
// @Tags Encrypt
// @Accept  json
// @Produce  json
// @Param msg query string false "Plaintext"
// @Param priv body Keypair true "Key"
// @Router /v1/encrypt2 [post]
// @Success 201 {object} response.Response
// @Failure 400,404,500 {object} response.ErrorResponse
func (data *EncryptHandler) Encrypt2(c *gin.Context) {
	msg := c.Query("msg")
	var input Keypair
	err := c.BindJSON(&input)
	if err != nil {
		code, result := response.BadRequest()
		c.JSON(code, result)
		return
	}

	var publicKey ecies.PublicKey
	var priv ecies.PrivateKey
	curveParams := &elliptic.CurveParams{Name: input.Curve.Name}
	curveParams.P, _ = new(big.Int).SetString(input.Curve.P, 10)
	curveParams.B, _ = new(big.Int).SetString(input.Curve.B, 10)
	curveParams.Gx, _ = new(big.Int).SetString(input.Curve.Gx, 10)
	curveParams.Gy, _ = new(big.Int).SetString(input.Curve.Gy, 10)
	curveParams.N, _ = new(big.Int).SetString(input.Curve.N, 10)
	curveParams.BitSize = 112

	publicKey.Curve = CurveWithParams{curveParams}

	publicKey.X, _ = new(big.Int).SetString(input.X, 10)
	publicKey.Y, _ = new(big.Int).SetString(input.Y, 10)

	priv.PublicKey = &publicKey
	priv.D, _ = new(big.Int).SetString(input.D, 10)

	ciphertext, err := ecies.Encrypt(priv.PublicKey, []byte(msg))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, ciphertext)
}
