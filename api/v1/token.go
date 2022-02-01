package v1

import (
	"net/http"
	"oauth_provider/utils/token"
	"oauth_provider/utils/verify"

	"github.com/gin-gonic/gin"
)

func CreateToken(c *gin.Context) {
	clientSecret := c.Query("client_secret")
	clientID := c.Query("client_id")
	ressource := c.Query("ressource")

	credentials, credentialsErr := verify.CredentialsExists(clientSecret, clientID)
	scope, scopeErr := verify.ScopeExists(ressource, credentials.ScopeIDs)

	if verified := verify.Client(scope.ID, credentials.ScopeIDs); verified != true || credentialsErr != nil || scopeErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorized"})
		return
	}

	tokenString := token.CreateJwt(*scope) // TODO: Remove password from here
	signedToken, err := token.Sign(tokenString)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": signedToken})
}
