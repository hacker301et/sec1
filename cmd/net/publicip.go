package net

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var getMyPublicIpCmd = &cobra.Command{
	Use:   "my-public-ip",
	Short: "allow you to get your current public ip",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("http://icanhazip.com")
		if err != nil {
			fmt.Println("unable to get your current public ip", err.Error())
			os.Exit(1)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("unable to get your public ip ", err.Error())
			os.Exit(1)
		}
		fmt.Printf("your current public ip => %s ", body)

	},
}

func init() {

}
