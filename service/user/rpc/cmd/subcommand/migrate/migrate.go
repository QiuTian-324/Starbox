package migrate

import (
	"BuzzBox/service/user/model/data"
	"BuzzBox/service/user/rpc/internal/config"
	"context"
	"flag"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"time"
)

var configFile = flag.String("migrateF", "D:\\AkitaCode\\Go\\project\\BuzzBox\\service\\user\\rpc\\etc\\user.yaml", "the config file")

// RunUp 执行数据库迁移操作
func RunUp(cmd *cobra.Command, args []string) {
	flag.Parse()

	var c config.Config
	// 加载配置文件
	conf.MustLoad(*configFile, &c)

	var mCtx context.Context
	var cancel context.CancelFunc

	// 获取迁移超时时间
	timeout, _ := cmd.Flags().GetUint("timeout")
	if timeout > 0 {
		// 设置上下文超时
		mCtx, cancel = context.WithTimeout(context.TODO(), time.Second*time.Duration(timeout))
		defer cancel()
	} else {
		mCtx = context.Background()
	}

	// 初始化数据库客户端
	cli, cleanFun, err := data.InitClientEnt(c.DB.Name, c.DB.DataSource)
	if err != nil {
		panic(err)
	} else {
		defer cleanFun()
	}

	// 获取调试标志
	debug, _ := cmd.Flags().GetBool("log")
	// 获取是否删除列标志
	dropColumn, _ := cmd.Flags().GetBool("drop-column")
	// 获取是否删除索引标志
	dropIndex, _ := cmd.Flags().GetBool("drop-index")
	// 获取是否创建外键标志
	foreignKey, _ := cmd.Flags().GetBool("create-foreign-key")
	// 获取是否使用唯一ID标志
	uniqueID, _ := cmd.Flags().GetBool("unique-id")

	// 执行数据库迁移
	data.Migrate(mCtx, cli, &data.MigrateOptions{
		Debug:            debug,
		DropColumn:       dropColumn,
		DropIndex:        dropIndex,
		CreateForeignKey: foreignKey,
		UniqueID:         uniqueID,
	})
}
