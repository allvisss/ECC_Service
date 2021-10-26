package api

import (
	"github.com/allvisss/ECC_service/common/model"
	"github.com/allvisss/ECC_service/common/response"
	service "github.com/allvisss/ECC_service/services/encrypt"
	"github.com/gin-gonic/gin"
)

type EncryptHandler struct {
}

type Keypair struct {
	Curve struct {
		P       int64  `json:"P"`
		N       int64  `json:"N"`
		B       int64  `json:"B"`
		Gx      int64  `json:"Gx"`
		Gy      int64  `json:"Gy"`
		BitSize int64  `json:"BitSize"`
		Name    string `json:"Name"`
	}
	X int64 `json:"X"`
	Y int64 `json:"Y"`
	D int64 `json:"D"`
}

func NewEncryptHandler(r *gin.Engine) {
	handler := &EncryptHandler{}
	v1 := r.Group("/v1")
	{
		v1.POST("/encrypt", handler.Encrypt)
	}
}

//Encrypt
// @Summary Encrypt plaintext into ciphertex (using curve secp112r1 as demo)
// @Tags Encrypt
// @Accept  json
// @Produce  json
// @Param msg query string false "Plaintext"
// @Param priv body Keypair true "Keypair"
// @Router /v1/encrypt [post]
// @Success 201 {object} response.Response
// @Failure 400,404,500 {object} response.ErrorResponse
func (data *EncryptHandler) Encrypt(c *gin.Context) {
	msg := c.Param("msg")
	var priv model.PrivateKey
	err := c.BindJSON(&priv)
	if err != nil {
		code, result := response.BadRequest()
		c.JSON(code, result)
		return
	}
	code, result := service.EncryptService(msg, priv)
	c.JSON(code, result)
}
