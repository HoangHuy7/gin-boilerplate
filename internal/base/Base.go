package base

import (
	"monorepo/internal/base/routerx"
	"monorepo/internal/dto"
)

type Controller interface {
	Register(*routerx.Routerx)
	GetMetadata() *dto.Metadata
}
