package controller

import (
	"monorepo/internal/base"
	"monorepo/internal/logger"
	"monorepo/internal/server"

	"github.com/swaggest/openapi-go/openapi3"
)

func NewRouter(controllers []base.Controller, lg *logger.GoLogger) *server.Router {
	rt := &server.Router{
		OpenAPI: openapi3.NewReflector(),
		Logger:  lg,
	}

	rt.OpenAPI.SpecEns().
		Info.
		WithTitle("Swagger Example API").
		WithVersion("1.0.0").
		WithDescription("This is an example api for swagger example")
	for _, c := range controllers {
		rt.Controllers = append(rt.Controllers, server.APIGroup{Controller: c})
	}
	return rt
}
