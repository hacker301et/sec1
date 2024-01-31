package net

import (
	"github.com/spf13/cobra"
)

var NetCmd = &cobra.Command{
	Use:   "net",
	Short: "net command for commands that related to network",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
func addCommand(){
	NetCmd.AddCommand(scanCmd)
	NetCmd.AddCommand(getMyPublicIpCmd)
}
func init() {
	addCommand()
}
