package main

import (
	"context"
	"github.com/p1ass/openapi-playground/ogen/petstore"
)

type petstoreSecurityHandler struct {
}

var _ petstore.SecurityHandler = &petstoreSecurityHandler{}

func (h *petstoreSecurityHandler) HandleAPIKey(ctx context.Context, operationName string, t petstore.APIKey) (context.Context, error) {
	// TODO implement me
	panic("implement me")
}
