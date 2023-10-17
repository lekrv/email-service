package main

import (
	"email-service/cmd/di"
	"email-service/internal/infrastructure/server"
	"email-service/pkg/mtools/mconfig"
	"log"
	"net/http"
)

func main() {
	mconfig.NewEnv()
	handler, err := di.Initialize()
	if err != nil {
		panic("fatal error:" + err.Error())
	}

	context := server.Context{Email: handler}
	serverInstance := server.NewRouter(&context)
	log.Fatal(http.ListenAndServe("localhost:8080", serverInstance.Router()))
}
