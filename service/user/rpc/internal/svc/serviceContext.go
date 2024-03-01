package svc

import (
	"BuzzBox/service/user/model/data"
	"BuzzBox/service/user/model/data/ent"
	"BuzzBox/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *ent.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	cli, _, _ := data.InitClientEnt(c.DB.Name, c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		DB:     cli,
	}
}
