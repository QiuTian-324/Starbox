package data

import (
	"BuzzBox/service/user/model/data/ent"
	"BuzzBox/service/user/model/data/ent/migrate"
	"context"
	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logc"
)

func InitClientEnt(DBType, DataSource string) *ent.Client {
	ctx := context.Background()

	// 打开数据库连接
	db, err := sql.Open(DBType, DataSource)
	if err != nil {
		panic("连接数据库失败:" + err.Error())
	}
	logc.Info(ctx, "数据库连接成功")
	// 程序退出关闭数据库连接
	defer func(drv *sql.Driver) {
		err := drv.Close()
		if err != nil {
			logc.Info(ctx, "数据库连接关闭失败:"+err.Error())
		}
	}(db)
	drv := ent.NewClient(ent.Driver(db))

	// 迁移数据库表
	err = drv.Debug().Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		logc.Info(ctx, "数据库表创建失败:"+err.Error())
	}

	logc.Info(ctx, "数据库表迁移成功")

	return drv
}
