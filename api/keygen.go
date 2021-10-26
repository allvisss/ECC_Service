package api

import (
	service "github.com/allvisss/ECC_service/services/keygen"
	"github.com/gin-gonic/gin"
)

type KeyGenerateHandler struct {
}

func NewKeyGenerateHandler(r *gin.Engine) {
	handler := &KeyGenerateHandler{}
	v1 := r.Group("/v1")
	{
		v1.GET("/keygen", handler.GetGeneratedKey)
	}
}

//Key Generation
// @Summary Generate key pair (using curve secp112r1 as demo)
// @Tags Key
// @Accept  */*
// @Produce  json
// @Router /v1/keygen [get]
// @Success 201 {object} response.Response
// @Failure 400,404,500 {object} response.ErrorResponse
func (data *KeyGenerateHandler) GetGeneratedKey(c *gin.Context) {
	code, result := service.GenerateKey()
	c.JSON(code, result)
}
