package service

import (
	"monorepo/apps/gas/app/database"
	"monorepo/internal/logger"
	"monorepo/shares/entities/workerdb"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CustomerService struct {
	db     *database.DataSources
	logger *zap.Logger
}

func NewProductService(db *database.DataSources, lg *logger.GoLogger) *CustomerService {
	return &CustomerService{
		db:     db,
		logger: lg.Zap,
	}
}

func (this *CustomerService) DeleteCustomer(id string) int64 {

	return 1
}

func (this *CustomerService) EditCustomer(c *gin.Context, input *workerdb.Gastb_Customer) (int64, error) {
	return 0, nil
}
