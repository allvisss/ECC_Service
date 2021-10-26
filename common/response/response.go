package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
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

func ServiceUnavailable() (int, interface{}) {
	return http.StatusServiceUnavailable, gin.H{
		"error": http.StatusText(http.StatusServiceUnavailable),
	}
}

func ServiceUnavailableMsg(msg interface{}) (int, interface{}) {
	return http.StatusServiceUnavailable, gin.H{
		"error": msg,
	}
}

func BadRequest() (int, interface{}) {
	return http.StatusBadRequest, gin.H{
		"error": http.StatusText(http.StatusBadRequest),
	}
}

func BadRequestMsg(msg interface{}) (int, interface{}) {
	return http.StatusBadRequest, gin.H{
		"error": msg,
	}
}

func NotFound() (int, interface{}) {
	return http.StatusNotFound, gin.H{
		"error": http.StatusText(http.StatusNotFound),
	}
}

func NotFoundMsg(msg interface{}) (int, interface{}) {
	return http.StatusNotFound, gin.H{
		"error": msg,
	}
}

func Forbidden() (int, interface{}) {
	return http.StatusForbidden, gin.H{
		"error": http.StatusText(http.StatusForbidden),
	}
}

func Unauthorized() (int, interface{}) {
	return http.StatusUnauthorized, gin.H{
		"error": http.StatusText(http.StatusUnauthorized),
	}
}
