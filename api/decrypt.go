package api

import (
	"crypto/elliptic"
	"math/big"
	"net/http"

	"github.com/allvisss/ECC_service/common/model"
	"github.com/allvisss/ECC_service/common/response"
	service "github.com/allvisss/ECC_service/services/decrypt"
	ecies "github.com/ecies/go"
	"github.com/gin-gonic/gin"
)

type DecryptHandler struct {
}

type CiphertextBlock struct {
	Ciphtext string `json:"Ciphertext"`
}

func NewDecryptHandler(r *gin.Engine) {
	handler := &DecryptHandler{}
	v1 := r.Group("/v1")
	{
		v1.POST("/decrypt", handler.Decrypt)
		v1.POST("/decrypt2", handler.Decrypt2)
	}
}

//Decrypt
// @Summary Decrypt ciphertext into plaintext (using curve secp112r1 as demo)
// @Tags Decrypt
// @Accept  json
// @Produce  json
// @Param ciphertext query string false "Ciphertext"
// @Param priv body Keypair true "Keypair"
// @Router /v1/decrypt [post]
// @Success 201 {object} response.Response
// @Failure 400,404,500 {object} response.ErrorResponse
func (data *DecryptHandler) Decrypt(c *gin.Context) {
	ciphertext := c.Query("ciphertext")
	var priv Keypair
	err := c.BindJSON(&priv)
	if err != nil {
		code, result := response.BadRequest()
		c.JSON(code, result)
		return
	}

	var publicKey model.PublicKey
	var privateKey model.PrivateKey
	curveParams := &elliptic.CurveParams{Name: priv.Curve.Name}
	curveParams.P, _ = new(big.Int).SetString(priv.Curve.P, 10)
	curveParams.B, _ = new(big.Int).SetString(priv.Curve.B, 10)
	curveParams.Gx, _ = new(big.Int).SetString(priv.Curve.Gx, 10)
	curveParams.Gy, _ = new(big.Int).SetString(priv.Curve.Gy, 10)
	curveParams.N, _ = new(big.Int).SetString(priv.Curve.N, 10)
	curveParams.BitSize = 112

	publicKey.Curve = CurveWithParams{curveParams}

	publicKey.X, _ = new(big.Int).SetString(priv.X, 10)
	publicKey.Y, _ = new(big.Int).SetString(priv.Y, 10)

	privateKey.PublicKey = &publicKey
	privateKey.D, _ = new(big.Int).SetString(priv.D, 10)

	code, result := service.DecryptService(ciphertext, privateKey)
	c.JSON(code, result)
}

//Decrypt2
// @Summary Decrypt ciphertext into plaintext
// @Tags Decrypt
// @Accept  json
// @Produce  json
// @Param ciphertext query string false "Ciphertext"
// @Param priv body Keypair true "Keypair"
// @Router /v1/decrypt2 [post]
// @Success 201 {object} response.Response
// @Failure 400,404,500 {object} response.ErrorResponse
func (data *DecryptHandler) Decrypt2(c *gin.Context) {
	ciphertext := c.Query("ciphertext")
	var priv Keypair
	err := c.BindJSON(&priv)
	if err != nil {
		code, result := response.BadRequest()
		c.JSON(code, result)
		return
	}

	var publicKey ecies.PublicKey
	var privateKey ecies.PrivateKey
	curveParams := &elliptic.CurveParams{Name: priv.Curve.Name}
	curveParams.P, _ = new(big.Int).SetString(priv.Curve.P, 10)
	curveParams.B, _ = new(big.Int).SetString(priv.Curve.B, 10)
	curveParams.Gx, _ = new(big.Int).SetString(priv.Curve.Gx, 10)
	curveParams.Gy, _ = new(big.Int).SetString(priv.Curve.Gy, 10)
	curveParams.N, _ = new(big.Int).SetString(priv.Curve.N, 10)
	curveParams.BitSize = 112

	publicKey.Curve = CurveWithParams{curveParams}

	publicKey.X, _ = new(big.Int).SetString(priv.X, 10)
	publicKey.Y, _ = new(big.Int).SetString(priv.Y, 10)

	privateKey.PublicKey = &publicKey
	privateKey.D, _ = new(big.Int).SetString(priv.D, 10)

	plaintext, err := ecies.Decrypt(&privateKey, []byte(ciphertext))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, plaintext)
}
