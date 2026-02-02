package service

import (
	"monorepo/apps/gas/app/database"
	"monorepo/internal/logger"
	"monorepo/internal/utils"
	"monorepo/shares/entities/workerdb"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
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

func (this *CustomerService) GetCustomerList() []workerdb.Gastb_Customer {
	list := make([]workerdb.Gastb_Customer, 0)
	tx := this.db.Worker.Order("id").Find(&list)
	if tx.Error != nil {
		panic(tx.Error)
	}
	return list
}

func (this *CustomerService) AddCustomer(customer *workerdb.Gastb_Customer) int64 {
	tx := this.db.Worker.Create(customer)
	if tx.Error != nil {
		panic(tx.Error)
	}
	return tx.RowsAffected
}

func (this *CustomerService) DeleteCustomer(id string) int64 {
	tx := this.db.Worker.Delete(&workerdb.Gastb_Customer{}, "id = ?", id)
	if tx.Error != nil {
		panic(tx.Error)
	}
	return tx.RowsAffected
}

func (this *CustomerService) EditCustomer(c *gin.Context, input *workerdb.Gastb_Customer) (int64, error) {

	var rows_affected int64

	err := this.db.Worker.Transaction(func(tx *gorm.DB) error {
		var customer_db workerdb.Gastb_Customer
		result := tx.First(&customer_db, "id = ?", input.Id)
		if er := utils.LogAndReturn(result.Error, this.logger); er != nil {
			return er
		}

		customer_db.Name = input.Name
		customer_db.Note = input.Note
		result = tx.Save(&customer_db)
		if er := utils.LogAndReturn(result.Error, this.logger); er != nil {
			return er
		}
		rows_affected = result.RowsAffected
		return nil
	})
	return rows_affected, err
}
