package main

import (
	"deploy/cmd/interactive"
	"deploy/version"

	"github.com/spf13/cobra"
)

var showVersion bool

func init() {
	RootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "show the version and exit")
	interactive.RegistrySubCommand(RootCmd)
}

var RootCmd = &cobra.Command{
	Use:   "deploy",
	Short: "`deploy`",
	Long:  "`deploy`",
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			version.PrintVersion()
			return
		}
		cmd.Usage()
	},
}

func main() {
	RootCmd.Execute()
}
