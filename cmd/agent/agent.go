package agent

import (
	"context"
	"os"

	agtMaster "github.com/blueseller/deploy.git/agent/chief"
	agtReplica "github.com/blueseller/deploy.git/agent/replica"
	"github.com/blueseller/deploy.git/cmd/common"
	"github.com/blueseller/deploy.git/dcontext"
	"github.com/blueseller/deploy.git/logger"
	"github.com/spf13/cobra"
)

func init() {
	agentMasterCmd.Flags().StringP("port", "p", ":54321", "agent master listen port")
	agentMasterCmd.Flags().StringP("ip", "i", "", "agent master listen port")

	agentCmd.Flags().StringP("address", "a", "", "please input agnet master addr")
	agentCmd.MarkFlagRequired("addr")
}

var agentCmd = &cobra.Command{
	Use:    "agent",
	Short:  "agent",
	Long:   "agent",
	PreRun: common.PreCmdRun,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		addr, err := cmd.Flags().GetString("addr")
		if err != nil {
			logger.GetContextLogger(ctx).Fatalf("input params addr is error. errmsg:%s", err.Error())
		}

		err = agtReplica.Run(ctx, addr)
		if err != nil {
		}
	},
}

var agentMasterCmd = &cobra.Command{
	Use:    "agent-master",
	Short:  "agent-master",
	Long:   "agent-master",
	PreRun: common.PreCmdRun,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := dcontext.GetDContext()

		// 输出agent 版本
		ip, _ := cmd.Flags().GetString("ip")
		port, _ := cmd.Flags().GetString("port")
		if port == "" && os.Getenv("DEPLOY_AGENT_MASTER_PORT") != "" {
			port = os.Getenv("DEPLOY_AGENT_MASTER_PORT")
		}
		ctx = dcontext.WithAgentMasterIp(ctx, ip)
		ctx = dcontext.WithAgentMasterPort(ctx, port)

		// 启动agent master 服务
		err := agtMaster.Run(ctx)
		if err != nil {
		}
	},
}

func RegistrySubCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(agentCmd)
	rootCmd.AddCommand(agentMasterCmd)
}
