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

var UpStep configure.CmdStep = 9999

type InputNum int

type CmdFlow struct {
	headerStep configure.CmdStep
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
func (c *CmdFlow) InitCmd(ctx context.Context, cmdFlowStruct configure.CmdFlow) (context.Context, error) {
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
	ctx = dcontext.WithCommandStep(ctx, configure.CmdStep(flowStepSort[0]))
	c.cmdStep = cmdFlowStruct
	c.headerStep = configure.CmdStep(flowStepSort[0])

	logger.GetContextLogger(ctx).Tracef("get cmd step map is %v", c)

	return ctx, nil
}

// 获取可执行的命令行信息
func (c *CmdFlow) GetWorkflowCmd(ctx context.Context) map[InputNum]interface{} {
	forwardStep := dcontext.CommandLastStep(ctx)
	if forwardStep == 0 {
		logger.GetContextLogger(ctx).Fatalf("have not interactive cmd")
		return nil
	}
	logger.GetContextLogger(ctx).Tracef("get work flow cmd step %d", forwardStep)

	nextSteps, thisCommands := c.getStepCommands(forwardStep)

	return c.printCmdGuide(nextSteps, thisCommands, forwardStep)
}

func doCommand(cmd configure.Command) {
	fmt.Println("开始运行 command", cmd.Name, cmd.Desc)
}

func (c *CmdFlow) printCmdGuide(nextSteps []configure.CmdStep, thisCommands []configure.Command, thisStep configure.CmdStep) (selects map[InputNum]interface{}) {
	selects = make(map[InputNum]interface{})
	var nextNum InputNum = 1
	for _, val := range nextSteps {
		if stepVar, ok := c.cmdStep[val]; ok {
			selects[nextNum] = stepVar
			fmt.Printf("%d. %s\n", nextNum, stepVar.Desc)
			nextNum++
		}
	}

	for _, val := range thisCommands {
		selects[nextNum] = val
		fmt.Printf("%d. %s\n", nextNum, val.Desc)
		nextNum++
	}

	if c.headerStep != thisStep {
		fmt.Printf("%d. 返回上一级\n", nextNum)
		selects[nextNum] = UpStep
	}
	fmt.Println("请选择一个标签进行操作?")
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
func (c *CmdFlow) WaitStdin(correctInput map[InputNum]interface{}) InputNum {
	reader := bufio.NewReader(os.Stdin)
	var nextNum InputNum
	for {
		data, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("!!!输入异常,请检查您的输入.")
			continue
		}

		// check data is correct
		num, err := strconv.Atoi(string(data))
		if err != nil {
			fmt.Println("!!!输入异常,请输入下一步编号.")
			continue
		}

		nextNum = InputNum(num)
		if _, ok := correctInput[nextNum]; !ok {
			fmt.Println("!!!输入异常,请输入上面展示的编号信息.")
			continue
		}

		break
	}
	return nextNum
}

// 根据用户的输入, 确定下一步的走向
func (c *CmdFlow) ExecInput(ctx context.Context, selectList map[InputNum]interface{}, next InputNum) context.Context {
	val := selectList[next]
	switch val.(type) {
	case configure.FlowCommand:
		logger.GetContextLogger(ctx).Tracef("选择了一个新的下级, %d", val.(configure.FlowCommand).Num)
		ctx = dcontext.WithCommandStep(ctx, val.(configure.FlowCommand).Num)
	case configure.Command:
		logger.GetContextLogger(ctx).Tracef("选择了一个可执行的命令, %d", val.(configure.Command).Hander)
		ctx = dcontext.WithCommand(ctx, val)
	case configure.CmdStep:
		logger.GetContextLogger(ctx).Tracef("选择了返回上级")
		inputHistory := dcontext.CommandStep(ctx)
		if len(inputHistory) >= 2 {
			ctx = dcontext.WithCommandStep(ctx, inputHistory[len(inputHistory)-2])
		}
	}
	return ctx
}

func (c *CmdFlow) DoHander(ctx context.Context) error {
	command := dcontext.Command(ctx)
	if command.Hander != "" {
		fmt.Println("do this command")
	}
	return nil
}
