package server

import (
	"fmt"
	"monorepo/internal/base"
	"monorepo/internal/base/routerx"
	"monorepo/internal/dto"
	"monorepo/internal/logger"
	"regexp"

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

func ginPathToOpenAPI(path string) string {
	re := regexp.MustCompile(`:([a-zA-Z0-9_]+)`)
	return re.ReplaceAllString(path, `{$1}`)
}

func (this *Router) registerOpenAPI(meta *dto.Metadata, base string) {

	for _, ep := range meta.Endpoints {
		// 1️⃣ tạo operation context
		oc, err := this.OpenAPI.NewOperationContext(
			ep.Method,
			base+ginPathToOpenAPI(ep.Path),
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
			"Gin Quickstathis API", // title
			"/openapi.json",        // OpenAPI spec URL
			"/swagger/",            // base path (BẮT BUỘC có / cuối)
		),
	))
}

func (this *Router) RegisterAll(r *gin.Engine) {
	for _, ctrl := range this.Controllers {
		meta := ctrl.Controller.GetMetadata()
		base := builderPath(meta)
		ginGroup := r.Group(base)
		rg := routerx.NewRouterx(ginGroup, meta)
		ctrl.Controller.Register(rg)
		if meta.EnableOpenAPI {
			this.registerOpenAPI(meta, base)
		}
		this.Logger.INFO("Registered controller",
			zap.String("url", base))
		//ctrl.Controller.Register(r.Group(builderPath(ctrl.Controller.GetMetadata())))
	}
	this.exposeSwagger(r)

}

func builderPath(config *dto.Metadata) string {
	return fmt.Sprintf("/api%s%s", config.Version, config.Path)
}
