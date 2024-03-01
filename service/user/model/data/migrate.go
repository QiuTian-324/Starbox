package data

import (
	"BuzzBox/service/user/model/data/ent"
	"BuzzBox/service/user/model/data/ent/migrate"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
)

// MigrateOptions 定义了数据库迁移的选项。
type MigrateOptions struct {
	Debug            bool // 是否启用调试模式。
	DropColumn       bool // 是否删除列。
	DropIndex        bool // 是否删除索引。
	CreateForeignKey bool // 是否创建外键。
	UniqueID         bool // 是否使用唯一ID。
}

func Migrate(ctx context.Context, cli *ent.Client, opt *MigrateOptions) {

	var err error

	if opt.Debug {
		// 迁移数据库表
		err = cli.Debug().Schema.Create(
			ctx,
			migrate.WithForeignKeys(opt.CreateForeignKey),
			migrate.WithDropIndex(opt.DropIndex),
			migrate.WithDropColumn(opt.DropColumn),
			migrate.WithGlobalUniqueID(opt.UniqueID),
		)
	} else {
		// 迁移数据库表
		err = cli.Schema.Create(
			ctx,
			migrate.WithForeignKeys(opt.CreateForeignKey),
			migrate.WithDropIndex(opt.DropIndex),
			migrate.WithDropColumn(opt.DropColumn),
			migrate.WithGlobalUniqueID(opt.UniqueID),
		)
	}

	if err != nil {
		panic(fmt.Sprintf("数据库表迁移失败:" + err.Error()))
	}

	logc.Info(ctx, "数据库表迁移成功")
}
