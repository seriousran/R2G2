package main

import (
  "github.com/savsgio/atreugo"
)

func main() {
	config := &atreugo.Config{
		Addr: "0.0.0.0:8000",
	}
	server := atreugo.New(config)

	// Register a route
	server.Path("GET", "/", func(ctx *atreugo.RequestCtx) error {
		return ctx.TextResponse("Hello R2G2")
	})

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}