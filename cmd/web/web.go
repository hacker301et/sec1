package web

import (
	"github.com/spf13/cobra"
)

var WebCmd = &cobra.Command{
	Use:   "web",
	Short: "website related tools",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func addSubCommands() {
	WebCmd.AddCommand(liveSubCmd)

}

func init() {
	addSubCommands()
}
