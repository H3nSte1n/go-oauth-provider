package main

import (
	"fmt"
	"oauth_provider/api"
)

func main() {
	// // Start the OAuth provider
	// oauth := NewOAuthProvider()
	// oauth.Start()

	// Start the API
	api.Init()
	fmt.Println("Service running!")
}
