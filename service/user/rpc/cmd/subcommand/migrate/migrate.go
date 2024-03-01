package migrate

import (
	"BuzzBox/service/user/model/data"
	"BuzzBox/service/user/rpc/internal/config"
	"context"
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"time"
)

var configFile = flag.String("migrateF", "D:\\AkitaCode\\Go\\project\\BuzzBox\\service\\user\\rpc\\etc\\user.yaml", "the config file")

// RunUp 数据库schema同步及数据初始化
func RunUp(cmd *cobra.Command, args []string) {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	var mCtx context.Context
	var cancel context.CancelFunc
	timeout, _ := cmd.Flags().GetUint("timeout")
	fmt.Println(timeout)
	if timeout > 0 {
		mCtx, cancel = context.WithTimeout(context.TODO(), time.Second*time.Duration(timeout))
		defer cancel()
	} else {
		mCtx = context.Background()
	}

	cli, cleanFun, err := data.InitClientEnt(c.DB.Name, c.DB.DataSource)
	if err != nil {
		panic(err)
	} else {
		defer cleanFun()
	}

	debug, _ := cmd.Flags().GetBool("log")
	dropColumn, _ := cmd.Flags().GetBool("drop-column")
	dropIndex, _ := cmd.Flags().GetBool("drop-index")
	foreignKey, _ := cmd.Flags().GetBool("create-foreign-key")
	uniqueID, _ := cmd.Flags().GetBool("unique-id")

	data.Migrate(mCtx, cli, &data.MigrateOptions{
		Debug:            debug,
		DropColumn:       dropColumn,
		DropIndex:        dropIndex,
		CreateForeignKey: foreignKey,
		UniqueID:         uniqueID,
	})
}
