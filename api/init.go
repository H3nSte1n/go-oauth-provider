package api

import (
	"oauth_provider/api/v1"
	"github.com/gin-gonic/gin"
	"fmt"
)

func Init() {
	r := gin.Default()
	r.GET("/credential", v1.CreateCredential)
	r.GET("/credentials", v1.GetCredentials)

	r.Run(":5002")
	fmt.Println("Service running!")
}
