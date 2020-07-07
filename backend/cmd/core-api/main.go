package main

import (
	"fmt"

	"github.com/awanku/awanku/internal/coreapi"
)

// @title Awanku API
// @version 0.1

// @contact.name Awanku Support
// @contact.email hello@awanku.id

// @host api.awanku.id
// @schemes	https

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://api.awanku.id/v1/auth/token
// @authorizationUrl https://api.awanku.id/v1/auth/{provider}/connect

func main() {
	fmt.Println("Starting server...")

	srv := coreapi.Server{}
	srv.Init()
	srv.Start()
}
