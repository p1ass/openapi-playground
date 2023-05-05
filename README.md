# OpenAPI Playground

## Generators

- [ogen](https://github.com/ogen-go/ogen)
- [oapi-codegen](https://github.com/deepmap/oapi-codegen)

## API Documents

https://redocly.github.io/redoc/?url=https://raw.githubusercontent.com/p1ass/openapi-playground/main/openapi.yaml

## How to Generate

```shell
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
go generate ./...
```

## Run

```shell
go run ./ogen/cmd/server
go run ./oapi-codegen/cmd/chi_server
go run ./oapi-codegen/cmd/echo_server
```

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=ogen-go/ogen,deepmap/oapi-codegen&type=Timeline)](https://star-history.com/#ogen-go/ogen&deepmap/oapi-codegen&Timeline)
