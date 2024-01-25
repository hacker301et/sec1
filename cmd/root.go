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
	Long: `
	🔐 Strengthen your defenses with our command-line security toolkit
	   – the ultimate solution for network and website penetration testing.
	    Identify vulnerabilities, fortify your digital assets, and stay one step ahead of threats, all from the command line.
	
        🔧 **Features:**
	   - Network Penetration Testing
	   - Website Security Assessment
	   - Real-Time Insights
	   - User-Friendly Interface
	
	🌐 **Why Choose Us:**
	   - Unified Solution
	   - Scalable for All Skill Levels
	   - Continuous Updates

	`,
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
