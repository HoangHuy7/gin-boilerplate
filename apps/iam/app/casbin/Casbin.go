package Casbin

import (
	"gin-quickstart/internal/database"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

type Casbin struct {
	Enforcer   *casbin.Enforcer
	datasource *database.DataSources
}

func NewCasbin(ds *database.DataSources) *Casbin {
	return &Casbin{
		Enforcer:   InitCasbin(ds),
		datasource: ds,
	}
}

func InitCasbin(ds *database.DataSources) *casbin.Enforcer {
	//e, err := casbin.NewEnforcer("configs/casbinConfig/auth_model.conf", "configs/casbinConfig/auth_policy.csv")
	a, err := gormadapter.NewAdapterByDB(ds.Master) // db là *gorm.DB
	if err != nil {
		panic(err)
	}
	e, err := casbin.NewEnforcer("configs/casbinConfig/auth_model.conf", a)
	if err != nil {
		panic(err)
	}

	err = e.LoadPolicy()
	if err != nil {
		panic(err)
	}

	return e
}
