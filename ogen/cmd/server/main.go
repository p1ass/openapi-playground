package main

import (
	"github.com/p1ass/openapi-playground/ogen/petstore"
	"log"
	"net/http"
)

func main() {
	service := &petService{}
	securityHandler := &petstoreSecurityHandler{}
	srv, err := petstore.NewServer(service, securityHandler)
	if err != nil {
		log.Fatalln(err)
	}

	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
