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
func main() {
	fmt.Println("Starting server...")

	srv := coreapi.Server{}
	srv.Init()
	srv.Start()
}
