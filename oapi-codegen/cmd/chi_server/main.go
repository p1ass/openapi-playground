package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/p1ass/openapi-playground/oapi-codegen/petstore_chi"
	"log"
	"net/http"
)

func main() {
	var service petService

	r := chi.NewRouter()
	r.Mount("/", petstore_chi.Handler(&service))

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
