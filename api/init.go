package api

import (
	"oauth_provider/api/v1"
	"github.com/gin-gonic/gin"
	"fmt"
)

func Init() {
	r := gin.Default()
	r.POST("/credentials", v1.CreateCredentials)
	r.GET("/credentials", v1.GetCredentials)

	r.POST("/service", v1.CreateService)
	r.PATCH("/service/:id", v1.UpdateService)
	r.GET("/service/:id", v1.GetService)
	r.GET("/services", v1.GetServices)


	r.Run(":5002")
	fmt.Println("Service running!")
}
