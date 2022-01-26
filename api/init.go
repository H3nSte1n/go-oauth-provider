package api

import (
	"github.com/gin-gonic/gin"
	"fmt"

	"oauth_provider/api/v1"
	"oauth_provider/middleware"
)

func Init() {
	r := gin.Default()
	r.POST("/credentials", middleware.TokenAuth(), v1.CreateCredentials)
	r.GET("/credentials", middleware.TokenAuth(), v1.GetCredentials)

	r.POST("/ressource", middleware.TokenAuth(), v1.CreateRessource)
	r.PATCH("/ressource/:id", middleware.TokenAuth(), v1.UpdateRessource)
	r.GET("/ressource/:id", middleware.TokenAuth(), v1.GetRessource)
	r.GET("/ressources", middleware.TokenAuth(), v1.GetRessources)

	r.POST("/user", middleware.TokenAuth(), v1.CreateUser)
	r.PATCH("/user/:id", middleware.TokenAuth(), v1.UpdateUser)
	r.DELETE("/user/:id", middleware.TokenAuth(), v1.DeleteUser)
	r.GET("/user/:id", middleware.TokenAuth(), v1.GetUser)
	r.GET("/users", middleware.TokenAuth(), v1.GetUsers)

	r.POST("/login", v1.Login)

	r.GET("/token", middleware.TokenAuth(), v1.CreateToken)


	r.Run(":5002")
	fmt.Println("Service running!")
}
