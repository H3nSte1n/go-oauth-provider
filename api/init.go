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

	r.POST("/ressource", v1.CreateRessource)
	r.PATCH("/ressource/:id", v1.UpdateRessource)
	r.GET("/ressource/:id", v1.GetRessource)
	r.GET("/ressources", v1.GetRessources)

	r.GET("/token", v1.CreateToken)


	r.Run(":5002")
	fmt.Println("Service running!")
}
