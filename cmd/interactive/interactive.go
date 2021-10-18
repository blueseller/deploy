package interactive

import (
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

	},
}
