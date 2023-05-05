package main

import (
	"github.com/labstack/echo/v4"
	"github.com/p1ass/openapi-playground/oapi-codegen/petstore_echo"
	"log"
)

func main() {
	service := &petService{}
	e := echo.New()
	petstore_echo.RegisterHandlers(e, service)

	if err := e.Start(":8080"); err != nil {
		log.Fatalln(err)
	}
}
