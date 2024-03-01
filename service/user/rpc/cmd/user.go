package main

import (
	"BuzzBox/service/user/rpc/cmd/subcommand/migrate"
	"BuzzBox/service/user/rpc/cmd/subcommand/server"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {

	// rootCmd 表示根命令，当没有任何子命令被调用时使用
	var rootCmd = &cobra.Command{
		Use:   "user",
		Short: "用户模块.",
		Long: `用户模块.
		默认运行 "serve". 使用 "-h" 参数查看所有可用命令.`,
	}

	// CmdServe 表示 "serve" 子命令
	var CmdServe = &cobra.Command{
		Use:   "serve",
		Short: "启动服务",
		Long:  `启动用户服务。提供服务及相关接口`,
		Run:   server.StartServe,
		Args:  cobra.MaximumNArgs(1),
	}

	// CmdMigrate 表示 "migrate" 子命令
	var CmdMigrate = &cobra.Command{
		Use:   "migrate",
		Short: "数据迁移",
		Long:  `数据迁移。将user下的schema同步到数据库，并进行相关数据初始化工作`,
		Run:   migrate.RunUp,
		Args:  cobra.MaximumNArgs(1),
	}

	//  初始化子命令
	{
		CmdMigrate.PersistentFlags().UintP("timeout", "t", 0, "迁移执行超时时间，单位：秒。大于等于0的整数，等于0时，永不超时。")
		CmdMigrate.PersistentFlags().BoolP("log", "v", false, "显示详细日志，如：打印sql日志等。")
		CmdMigrate.PersistentFlags().Bool("drop-column", false, "是否删除原有字段")
		CmdMigrate.PersistentFlags().Bool("drop-index", false, "是否删除原有索引")
		CmdMigrate.PersistentFlags().Bool("foreign-key", true, "是否创建外键")
		CmdMigrate.PersistentFlags().Bool("unique-id", true, "是否使用唯一ID。")

	}

	//  将子命令添加到根命令中
	rootCmd.AddCommand(
		CmdServe,
		CmdMigrate,
	)
	// 执行命令
	if err := rootCmd.Execute(); err != nil {
		panic(fmt.Sprintf("运行应用程序失败 %v: %s", os.Args, err.Error()))
	}
}
