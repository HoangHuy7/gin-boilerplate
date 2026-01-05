package customer

import (
	"monorepo/apps/gas/service"
	"monorepo/internal/base/routerx"
	"monorepo/internal/dto"
	"monorepo/internal/utils"
	"monorepo/shares/entities/workerdb"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	Metadata        dto.Metadata
	customerService *service.CustomerService
}

func (p *CustomerController) GetMetadata() *dto.Metadata {
	return &p.Metadata
}
func NewCustomerController(ps *service.CustomerService) *CustomerController {
	return &CustomerController{
		customerService: ps,
		Metadata: dto.Metadata{
			Tag:           "Product Controller",
			Version:       "/v1",
			Path:          "/customer",
			Endpoints:     []dto.OpenEndpoint{},
			EnableOpenAPI: true,
			IsNotAuth:     true,
		},
	}
}
func (this *CustomerController) Register(rg *routerx.Routerx) {
	rg.GET(dto.OpenEndpoint{
		Path:        "",
		Handler:     this.filter,
		Summary:     "Hello World Summary",
		Description: "Hello World Description",
	})
	rg.POST(dto.OpenEndpoint{
		Path:        "",
		Handler:     this.add,
		Summary:     "Hello World Summary",
		Description: "Hello World Description",
		Request:     workerdb.Gastb_Customer{},
	})
	rg.DELETE(dto.OpenEndpoint{
		Path:        "",
		Handler:     this.remove,
		Summary:     "Hello World Summary",
		Description: "Hello World Description",
		Query: struct {
			ID string `form:"id"`
		}{},
	})
	rg.PUT(dto.OpenEndpoint{
		Path:        "",
		Handler:     this.edit,
		Summary:     "Hello World Summary",
		Description: "Hello World Description",
		Request:     workerdb.Gastb_Customer{},
	})
}
func (this CustomerController) edit(c *gin.Context) {

	var customer workerdb.Gastb_Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(400, utils.ErrorResponse(err))
		return
	}
	rows_affected, err := this.customerService.EditCustomer(&customer)
	if err != nil {
		c.JSON(500, utils.ErrorResponse(err))
		return
	}
	c.JSON(200, gin.H{
		"rows_affected": rows_affected,
	})

}
func (this CustomerController) remove(c *gin.Context) {

	id := c.Query("id")
	if id == "" {
		c.JSON(400, gin.H{
			"error": "id is required",
		})
		return
	}

	c.JSON(200, gin.H{
		"rows_affected": this.customerService.DeleteCustomer(id),
	})

}
func (this CustomerController) add(c *gin.Context) {

	var customer workerdb.Gastb_Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"rows_affected": this.customerService.AddCustomer(&customer),
	})
}
func (this *CustomerController) filter(c *gin.Context) {
	c.JSON(200, this.customerService.GetCustomerList())
}
