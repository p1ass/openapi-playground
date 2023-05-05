package main

import (
	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	"github.com/p1ass/openapi-playground/oapi-codegen/petstore_chi"
	"log"
	"net/http"
	"os"
)

func main() {
	swagger, err := petstore_chi.GetSwagger()
	if err != nil {
		log.Printf("Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	var service petService

	r := chi.NewRouter()
	r.Use(middleware.OapiRequestValidator(swagger))
	r.Mount("/", petstore_chi.Handler(&service))

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
