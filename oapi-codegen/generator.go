package oapi_codegen

//go:generate oapi-codegen --package petstore_echo -generate types,client,server,spec -o ./petstore_echo/petstore.gen.go ../openapi.yaml
//go:generate oapi-codegen --package petstore_chi -generate types,client,chi-server,spec -o ./petstore_chi/petstore.gen.go ../openapi.yaml
