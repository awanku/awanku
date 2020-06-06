package main

import (
	"net/http"

	"gopkg.in/macaron.v1"
)

func main() {
	m := macaron.Classic()
	m.Get("/api/", func() string {
		return "Nothing here"
	})
	http.ListenAndServe("0.0.0.0:3000", m)
}
