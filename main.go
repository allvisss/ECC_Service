package main

import (
	api "github.com/allvisss/ECC_service/api"
	"github.com/gin-gonic/gin"
)

func init() {

}

// @title ECC SERVICE API
// @version 0.0
// @description Elliptic Curve Cryptography service API for IOT system
// @termsOfService http://swagger.io/terms/

// @contact.name NGUYEN LE BAO TRUONG
// @contact.url https://github.com/allvisss
// @contact.email truongnlbse140940@fpt.edu.vn

// @host
// @BasePath
func main() {
	r := gin.Default()
	api.SwaggerHandler(r)
	api.NewTestHandler(r)

	r.Run("localhost:8080")
}
