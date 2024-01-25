package cmd

import (
	"os"

	"github.com/hacker301et/sec1/cmd/net"
	"github.com/hacker301et/sec1/cmd/web"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sec1",
	Short: "Command-Line Security Toolkit: For Network & Web Pen Testing",
	Long: ``,
}

func addSubCommands() {
	rootCmd.AddCommand(net.NetCmd)
	rootCmd.AddCommand(web.WebCmd)
}
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	addSubCommands()

}
