package interactive

import (
	"context"

	"github.com/blueseller/deploy.git/dcontext"
	"github.com/blueseller/deploy/cmd/interactive/flow"
	"github.com/blueseller/deploy/configure"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func RegistrySubCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(interactiveCmd)
}

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "interactive",
	Long:  "interactive",
	Run: func(cmd *cobra.Command, args []string) {
		//ctx := context.Background()

		// parse config file
		//	config, err := resolveConfiguration(args)
		//	if err != nil {
		//		logrus.Fatalf("%s", err.Error())
		//	}

		//	// 解析并设定 logger
		//	ctx, err = logger.LoggerFactory(ctx, config)
		//	if err != nil {
		//		logrus.Fatalf("%s", err.Error())
		//	}

		config := configure.GetConfig()
		ctx := dcontext.GetDContext()

		// 开始监听命令行输入
		StartInteractive(ctx, config)
	},
}

func StartInteractive(ctx context.Context, config *configure.Configuration) {
	flowSrv := flow.NewCmdFlow(config)
	var err error
	ctx, err = flowSrv.InitCmd(ctx, config.CmdFlow)
	if err != nil {
		logrus.Fatalf("init cmd flow data is error, %v", err)
	}
	for {
		// 获取现在可执行的命令
		selectList := flowSrv.GetWorkflowCmd(ctx)

		// 等待用户的输入
		next := flowSrv.WaitStdin(selectList)

		// 处理输入内容
		ctx = flowSrv.ExecInput(ctx, selectList, next)

		// 执行命令
		err = flowSrv.DoHander(ctx)
		if err != nil {
			logrus.Fatalf("exec input command is error, %v", err)
		}
	}
}
