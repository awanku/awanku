package main

import (
	"github.com/awanku/awanku/backend/internal/core"
)

func main() {
	srv := core.Server{}
	srv.Init()
	srv.Start()
}
