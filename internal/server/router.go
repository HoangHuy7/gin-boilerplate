// hoanghuy7 from Vietnamese with love!

package server

import (
	"fmt"
	"monorepo/internal/base"
	"monorepo/internal/base/routerx"
	"monorepo/internal/dto"
	"monorepo/internal/logger"
	"monorepo/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/swaggest/openapi-go"
	"github.com/swaggest/openapi-go/openapi3"
	"github.com/swaggest/swgui/v5emb"
	"go.uber.org/zap"
)

type APIGroup struct {
	Controller base.Controller
}

type Router struct {
	Controllers []APIGroup
	OpenAPI     *openapi3.Reflector
	Logger      *logger.GoLogger
}

func (this *Router) registerOpenAPI(meta *dto.Metadata, base string) {

	for _, ep := range meta.Endpoints {
		// 1️⃣ tạo operation context
		oc, err := this.OpenAPI.NewOperationContext(
			ep.Method,
			base+utils.GinPathToOpenAPI(ep.Path),
		)
		if err != nil {
			panic(err) // hoặc log
		}

		// 2️⃣ set summary (nếu có)
		if ep.Summary != "" {
			oc.SetSummary(ep.Summary)
		}
		if ep.Description != "" {
			oc.SetDescription(ep.Description)
		}
		if meta.Tag != "" {
			oc.SetTags(meta.Tag)
		}
		// 3️⃣ response schema (200)
		// 2️⃣.5 request body
		if ep.Request != nil {
			oc.AddReqStructure(
				ep.Request,
				func(cu *openapi.ContentUnit) {
					cu.ContentType = "application/json"
				},
			)
		}
		if ep.Query != nil {
			oc.AddReqStructure(ep.Query)
		}

		if ep.Responses != nil {
			for status, respSchema := range ep.Responses {
				oc.AddRespStructure(
					respSchema,
					func(cu *openapi.ContentUnit) {
						cu.ContentType = "application/json"
						cu.HTTPStatus = status
					},
				)
			}

		}

		// 4️⃣ add operation vào spec
		if err := this.OpenAPI.AddOperation(oc); err != nil {
			panic(err)
		}
	}
}
func (this *Router) exposeSwagger(r *gin.Engine) {

	r.GET("/openapi.json", func(c *gin.Context) {
		spec := this.OpenAPI.SpecEns()
		c.JSON(200, spec)
	})
	r.GET("/swagger/*any", gin.WrapH(
		v5emb.New(
			"Gin Quickstathis API",
			"/openapi.json",
			"/swagger/",
		),
	))
}

func (this *Router) RegisterAll(r *gin.Engine, appMeta *dto.AppMetadata) {
	for _, ctrl := range this.Controllers {
		meta := ctrl.Controller.GetMetadata()
		base := builderPath(meta, appMeta.ContextPath)
		ginGroup := r.Group(base)
		rg := routerx.NewRouterx(ginGroup, meta)
		ctrl.Controller.Register(rg)
		if this.OpenAPI != nil {
			if meta.EnableOpenAPI {
				this.registerOpenAPI(meta, base)
			}
			this.Logger.INFO("Registered controller",
				zap.String("url", base))
		} else {
			this.Logger.Zap.Info("OpenAPI is disabled; controller not registered in OpenAPI spec")
		}

		//ctrl.Controller.Register(r.Group(builderPath(ctrl.Controller.GetMetadata())))
	}
	if this.OpenAPI != nil {
		this.exposeSwagger(r)
	}

}

func builderPath(config *dto.Metadata, contextPath string) string {
	return fmt.Sprintf("%s/api%s%s", contextPath, config.Version, config.Path)
}

func NewRouter(controllers []base.Controller, appMeta *dto.AppMetadata, lg *logger.GoLogger) *Router {
	rt := &Router{
		OpenAPI: openapi3.NewReflector(),
		Logger:  lg,
	}

	rt.OpenAPI.SpecEns().
		Info.
		WithTitle(utils.NVL(&appMeta.OpenAPIInfo.Title, fmt.Sprintf("Swagger Service %s", appMeta.AppName))).
		WithVersion("1.0.0").
		WithDescription("This is an example api for swagger example")
	for _, c := range controllers {
		rt.Controllers = append(rt.Controllers, APIGroup{Controller: c})
	}
	return rt
}

func NewRouterWithOpenAPI(controllers []base.Controller, lg *logger.GoLogger) *Router {
	rt := &Router{
		OpenAPI: nil,
		Logger:  lg,
	}
	for _, c := range controllers {
		rt.Controllers = append(rt.Controllers, APIGroup{Controller: c})
	}
	return rt
}
