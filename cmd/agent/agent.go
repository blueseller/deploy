package agent

import (
	"context"
	"os"

	agtMaster "github.com/blueseller/deploy/agent/chief"
	"github.com/blueseller/deploy/dcontext"
	"github.com/spf13/cobra"
)

func init() {
	agentMasterCmd.Flags().StringP("port", "p", ":54321", "agent master listen port")
	agentMasterCmd.Flags().StringP("ip", "ip", "", "agent master listen port")
}

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "agent",
	Long:  "agent",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
	},
}

var agentMasterCmd = &cobra.Command{
	Use:   "agent-master",
	Short: "agent-master",
	Long:  "agent-master",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := dcontext.GetDContext()

		// 输出agent 版本
		ip := testCmd.Flags().GetString("ip")
		port := testCmd.Flags().GetString("port")
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
