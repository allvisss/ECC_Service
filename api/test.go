package api

import (
	"net/http"

	"github.com/allvisss/ECC_service/common/response"
	"github.com/gin-gonic/gin"

	ecies "github.com/ecies/go"
)

type TestHandler struct {
}

type DecryptResponse struct {
	Plaintext  string `json:"Plaintext"`
	Ciphertext string `json:"Ciphertext"`
}

func NewTestHandler(r *gin.Engine) {
	handler := &TestHandler{}
	v1 := r.Group("/v1")
	{
		v1.GET("/testDecrypt", handler.TestDecrypt)
	}
}

//Test
// @Summary Testing decrypt method
// @Tags Test
// @Accept  json
// @Produce  json
// @Param plaintext query string false "Plaintext"
// @Router /v1/testDecrypt [get]
// @Success 200 {object} response.Response
// @Failure 400,404,500 {object} response.ErrorResponse
func (data *TestHandler) TestDecrypt(c *gin.Context) {
	k, err := ecies.GenerateKey()
	plain := c.Query("plaintext")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	ciphertext, err := ecies.Encrypt(k.PublicKey, []byte(plain))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	plaintext, err := ecies.Decrypt(k, ciphertext)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	var dat DecryptResponse
	dat.Ciphertext = string(ciphertext)
	dat.Plaintext = string(plaintext)

	code, result := response.NewOKResponse(dat)
	c.JSON(code, result)
}
