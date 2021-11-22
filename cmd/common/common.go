package common

import (
	"context"
	"fmt"
	"os"

	"github.com/blueseller/deploy.git/configure"
	"github.com/blueseller/deploy.git/dcontext"
	"github.com/blueseller/deploy.git/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func PreCmdRun(cmd *cobra.Command, args []string) {
	// parse config file
	config, err := resolveConfiguration(args)
	if err != nil {
		logrus.Fatalf("%s", err.Error())
	}

	configure.SetConfig(config)
	//cfg := configure.GetConfig()

	// 解析并设定 tr := testCmd.Flags().GetString("aaa")str := testCmd.Flags().GetString("aaa")ogger
	ctx, err := logger.LoggerFactory(context.Background(), config)
	if err != nil {
		logrus.Fatalf("%s", err.Error())
	}

	dcontext.SetDContext(ctx)
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

//Deprecated
func initLoggerLevel(ctx context.Context, config *configure.Configuration) {
	level, err := logrus.ParseLevel(string(config.Log.LogLevel))
	if err != nil {
		level = logrus.InfoLevel
		logrus.Warnf("error parse log level %+s : %v, using %q", string(config.Log.LogLevel), err, level)
	}
	logrus.SetLevel(level)
}
