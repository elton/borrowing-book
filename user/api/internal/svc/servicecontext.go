package svc

import (
	"github.com/elton/borrowing-book/user/api/internal/config"
	"github.com/elton/borrowing-book/user/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel *model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.Datasource)
	um := model.NewUserModel(conn, c.CacheRedis)
	return &ServiceContext{
		Config:    c,
		UserModel: um,
	}
}
