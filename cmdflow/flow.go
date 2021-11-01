package cmdflow

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/blueseller/deploy/configure"
	"github.com/blueseller/deploy/dcontext"
	"github.com/blueseller/deploy/logger"
)

type CmdFlow struct {
	stepHeader configure.CmdStep
	cmdStep    map[configure.CmdStep]configure.FlowCommand
}

func NewCmdFlow(config *configure.Configuration) *CmdFlow {
	flow := new(CmdFlow)
	return flow
}

// 初始化命令执行流程
// 检查是否存在一些不能够完成执行命令的阶段
// 此阶段梳理整体命令行流程, 并输出一个完成路径图
// 排序严格按照 数据顺序
// next step 优先级高于 commands
func (c *CmdFlow) InitCmd(ctx context.Context, cmdFlowStruct configure.CmdFlow) error {
	// 排序准备
	flowStepSort := make([]int, 0)
	for step, _ := range cmdFlowStruct {
		flowStepSort = append(flowStepSort, int(step))

		cmds := cmdFlowStruct[step]
		nextSteps := cmds.NextCmdSteps
		if len(nextSteps) > 0 {
			// 检查next step 是否存在
			for _, next := range nextSteps {
				if _, ok := cmdFlowStruct[next]; !ok {
					logger.GetContextLogger(ctx).Fatalf("can not find next step int cmdflow, please check you yaml input")
				}
			}

			//TODO 检查命令路由是否存在，参数是否正确
		}
	}

	// sort
	sort.Ints(flowStepSort)
	c.stepHeader = configure.CmdStep(flowStepSort[0])
	c.cmdStep = cmdFlowStruct

	logger.GetContextLogger(ctx).Tracef("get cmd step map is %v", c)

	return nil
}

// 获取可执行的命令行信息
func (c *CmdFlow) GetWorkflowCmd(ctx context.Context) {
	thisStep := c.stepHeader

	forwardStep := dcontext.CommandUpStep(ctx)
	if int(forwardStep) > 0 {
		thisStep = forwardStep
	}
	logger.GetContextLogger(ctx).Tracef("get work flow cmd step %d", thisStep)

	guideCmds := c.getStepCommands(thisStep)
	printCmdGuide(guideCmds)

	// lience input
	c.waitStdin()
}

func printCmdGuide(guides []string) {
	for num, desc := range guides {
		fmt.Printf("%d. %s \n", num, desc)
	}
	fmt.Println("请选择一个标签进行操作")
}

func (c *CmdFlow) getStepCommands(step configure.CmdStep) (cmdList []string) {
	if stepVar, ok := c.cmdStep[step]; ok {
		cmdList = append(cmdList, c.getNextSteps(stepVar.NextCmdSteps)...)
		cmdList = append(cmdList, c.getNextCmds(stepVar.Commands)...)
	}
	return cmdList
}

func (c *CmdFlow) getNextSteps(list []configure.CmdStep) (cmdList []string) {
	for _, step := range list {
		if setepVal, ok := c.cmdStep[step]; ok {
			cmdList = append(cmdList, setepVal.Desc)
		}
	}
	return cmdList
}
func (c *CmdFlow) getNextCmds(list []configure.Command) (cmdList []string) {
	for _, v := range list {
		cmdList = append(cmdList, v.Desc)
	}
	return cmdList
}

// 等待输入, 检测输入的是否正确，如果不正确则重新等待输入
func (c *CmdFlow) waitStdin() {
	reader := bufio.NewReader(os.Stdin)
	inputSate := false
	for !inputSate {
		data, _, _ := reader.ReadLine()

	}
	fmt.Println("")
}
