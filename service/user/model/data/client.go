package data

import (
	"BuzzBox/service/user/model/data/ent"
	"context"
	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logc"
)

//var Cli *ent.Client

func InitClientEnt(DBType, DataSource string) (*ent.Client, func(), error) {
	ctx := context.Background()

	// 打开数据库连接
	drv, err := sql.Open(DBType, DataSource)
	if err != nil {
		panic("连接数据库失败:" + err.Error())
	}

	// 定义用于关闭资源的 cleanup 函数
	cleanup := func() {
		if err := drv.Close(); err != nil {
			logc.Info(ctx, "关闭数据库连接失败:"+err.Error())
		}
		logc.Info(ctx, "关闭数据库连接成功")
	}

	cli := ent.NewClient(ent.Driver(drv))
	logc.Info(ctx, "数据库连接成功")
	return cli, cleanup, nil
}
