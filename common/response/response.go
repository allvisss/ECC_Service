package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
}

func NewResponse(code int, data interface{}) (int, interface{}) {
	return code, gin.H{
		"data": data,
	}
}

func NewOKResponse(data interface{}) (int, interface{}) {
	return http.StatusOK, gin.H{
		"data": data,
	}
}
