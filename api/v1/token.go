package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oauth_provider/utils"
	"oauth_provider/utils/verify"
	"oauth_provider/models"
)

func CreateToken(c *gin.Context) {
	client_secret := c.Query("client_secret")
	client_id := c.Query("client_id")
	ressource := c.Query("ressource")

	credentials, credentials_err := verify.CredentialsExists(client_secret, client_id)
	scope, scope_err := verify.ScopeExists(ressource, credentials.ScopeIDs)
	if verified := verify.Client(scope.ID, credentials.ScopeIDs); verified != true || credentials_err != nil || scope_err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorized"})
		return
	}
	
	token := utils.CreateAccessToken[models.Scope](*scope) // TODO: Remove password from here
	signedToken, err := utils.SignToken(token)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": signedToken})
}