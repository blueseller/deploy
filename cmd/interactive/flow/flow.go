package flow

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/blueseller/deploy/configure"
	"github.com/blueseller/deploy/dcontext"
	"github.com/blueseller/deploy/logger"
)

type InputNum int

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

	nextSteps, thisCommands := c.getStepCommands(thisStep)

	selects := c.printCmdGuide(nextSteps, thisCommands)

	// lience input
	next := c.waitStdin(selects)
	nextNum := int(next)

	// 检查 用户操作的是 下一步目录 还是 一个命令
	if nextNum > len(nextSteps) {
		nextNum = nextNum - len(nextSteps)
		doCommand(thisCommands[nextNum-1])
	}

}

func doCommand(cmd configure.Command) {
	fmt.Println("开始运行 command", cmd.Name, cmd.Desc)
}

func (c *CmdFlow) printCmdGuide(nextSteps []configure.CmdStep, thisCommands []configure.Command) (selects map[InputNum]string) {
	selects = make(map[InputNum]string)
	var nextNum InputNum = 1
	for _, val := range nextSteps {
		if stepVar, ok := c.cmdStep[val]; ok {
			selects[nextNum] = stepVar.Desc
			fmt.Printf("%d. %s\n", nextNum, stepVar.Desc)
			nextNum++
		}
	}
	for _, val := range thisCommands {
		selects[nextNum] = val.Desc
		fmt.Printf("%d. %s\n", nextNum, val.Desc)
		nextNum++
	}
	fmt.Println("请选择一个标签进行操作")
	return
}

func (c *CmdFlow) getStepCommands(step configure.CmdStep) (nextSteps []configure.CmdStep, thisCommands []configure.Command) {
	if stepVar, ok := c.cmdStep[step]; ok {
		nextSteps = stepVar.NextCmdSteps
		thisCommands = stepVar.Commands
	}
	return
}

// 等待输入, 检测输入的是否正确，如果不正确则重新等待输入
func (c *CmdFlow) waitStdin(correctInput map[InputNum]string) InputNum {
	reader := bufio.NewReader(os.Stdin)
	var nextNum InputNum
	for {
		data, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("输入异常,请检查您的输入.")
			continue
		}

		// check data is correct
		num, err := strconv.Atoi(string(data))
		if err != nil {
			fmt.Println("输入异常,请输入下一步编号.")
			continue
		}

		nextNum = InputNum(num)
		if _, ok := correctInput[nextNum]; !ok {
			fmt.Println("输入异常,请输入上面展示的编号信息.")
			continue
		}

		break
	}
	return nextNum
}
