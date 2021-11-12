package agent

import (
	"context"

	"github.com/spf13/cobra"
)

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
		ctx := context.Background()

	},
}
