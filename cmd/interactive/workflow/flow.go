package workflow

import (
	"context"
)

func StartInteractive(ctx context.Context) {
	for {
		// 获取现在可执行的命令
		cmdStr := getWorkflowCmd(ctx)
	}
}

func getWorkflowCmd(ctx context.Context) {

}
