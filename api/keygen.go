package api

import (
	"net/http"

	service "github.com/allvisss/ECC_service/services/keygen"
	ecies "github.com/ecies/go"
	"github.com/gin-gonic/gin"
)

type KeyGenerateHandler struct {
}

func NewKeyGenerateHandler(r *gin.Engine) {
	handler := &KeyGenerateHandler{}
	v1 := r.Group("/v1")
	{
		v1.GET("/keygen", handler.GetGeneratedKey)
		v1.GET("/keygen2", handler.GetGeneratedKey2)
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

//Key Generation 2
// @Summary Generate key pair
// @Tags Key
// @Accept  */*
// @Produce  json
// @Router /v1/keygen2 [get]
// @Success 201 {object} response.Response
// @Failure 400,404,500 {object} response.ErrorResponse
func (data *KeyGenerateHandler) GetGeneratedKey2(c *gin.Context) {
	k, err := ecies.GenerateKey()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	err = service.WriteKeyToFile2(k)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, k)
}
