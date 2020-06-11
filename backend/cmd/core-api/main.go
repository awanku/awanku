package main

import "github.com/awanku/awanku/internal/coreapi"

func main() {
	srv := coreapi.Server{}
	srv.Init()
	srv.Start()
}
