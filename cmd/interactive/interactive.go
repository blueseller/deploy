package interactive

import (
	"context"
	"fmt"
	"os"

	"github.com/blueseller/deploy/cmd/interactive/workflow"
	"github.com/blueseller/deploy/configure"
	"github.com/blueseller/deploy/logger"

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
		ctx := context.Background()

		// parse config file
		config, err := resolveConfiguration(args)
		if err != nil {
			logrus.Fatalf("%s", err.Error())
		}

		// 解析并设定 logger
		ctx, err = logger.LoggerFactory(ctx, config)
		if err != nil {
			logrus.Fatalf("%s", err.Error())
		}

		// 开始监听命令行输入
		workflow.StartInteractive(ctx, config)
	},
}

func resolveConfiguration(args []string) (*configure.Configuration, error) {
	var configurationPath string

	if len(args) > 0 {
		configurationPath = args[0]
	} else if os.Getenv("DEPLOY_CONFIGURATION_PATH") != "" {
		configurationPath = os.Getenv("DEPLOY_CONFIGURATION_PATH")
	}

	if configurationPath == "" {
		return nil, fmt.Errorf("configuration path is unspecified")
	}

	fp, err := os.Open(configurationPath)
	if err != nil {
		return nil, err
	}

	defer fp.Close()

	config, err := configure.Parse(fp)
	if err != nil {
		return nil, fmt.Errorf("error parsing %s:%v", configurationPath, err)
	}

	return config, nil

}

func initLoggerLevel(ctx context.Context, config *configure.Configuration) {
	level, err := logrus.ParseLevel(string(config.Log.LogLevel))
	if err != nil {
		level = logrus.InfoLevel
		logrus.Warnf("error parse log level %+s : %v, using %q", string(config.Log.LogLevel), err, level)
	}
	logrus.SetLevel(level)
}
