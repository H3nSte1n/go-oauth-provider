package main

import (
	"fmt"
	"oauth_provider/api"
)

func main() {
	// Start the API
	api.Init()
	fmt.Println("Service running!")
}
