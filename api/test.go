package api

import (
	"github.com/allvisss/ECC_service/common/response"
	"github.com/gin-gonic/gin"
)

type TestHandler struct {
}

func NewTestHandler(r *gin.Engine) {
	handler := &TestHandler{}
	v1 := r.Group("/v1")
	{
		v1.GET("/test", handler.TestGET)
	}
}

//Test GET
// @Summary Testing API GET method
// @Tags Test
// @Accept  */*
// @Produce  json
// @Router /v1/test [get]
// @Success 200 {object} response.Response
// @Failure 400,404,500 {object} response.ErrorResponse
func (data *TestHandler) TestGET(c *gin.Context) {
	code, result := response.NewOKResponse("API Working")
	c.JSON(code, result)
}
