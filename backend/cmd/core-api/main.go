package main

import (
	"fmt"

	"github.com/awanku/awanku/internal/coreapi"
)

func main() {
	fmt.Println("Starting server...")

	srv := coreapi.Server{}
	srv.Init()
	srv.Start()
}
