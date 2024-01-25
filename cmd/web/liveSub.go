package web

import (
	livesub "github.com/hacker301et/sec1/cmd/web/subfinder_tool/cmd"
	"github.com/spf13/cobra"
)

var liveSubCmd = &cobra.Command{
	Use:   "live-sub",
	Short: "get live sub domains using different tools",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		livesub.Start()
	},
}

func init() {

}
