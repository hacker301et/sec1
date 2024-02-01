package net

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var iface string
var mac string

func generateRandomMAC() string {
	rand.Seed(time.Now().UnixNano())

	pairs := make([]string, 6)
	for i := 0; i < 6; i++ {
		pairs[i] = fmt.Sprintf("%02x", rand.Intn(256))
	}
	macAddress := fmt.Sprintf("%s:%s:%s:%s:%s:%s", pairs[0], pairs[1], pairs[2], pairs[3], pairs[4], pairs[5])
	return macAddress
}

var chageMacCmd = &cobra.Command{
	Use:   "change-mac",
	Short: "change-mac  is a utility to change MAC address of network interface.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(mac) != 11 {
			mac = generateRandomMAC()
		}
		command := exec.Command("bash", "-c", fmt.Sprintf("ifconfig %s | grep ether | awk '{print $2}'", iface))
		out, err := command.CombinedOutput()
		if err != nil {
			fmt.Println("Error : ", err)
			os.Exit(1)
		}
		fmt.Println(fmt.Sprintf("Current MAC  Address  =>  %s", strings.TrimSpace(string(out))))
		// Change Mac Addrss
		command = exec.Command("ifconfig", iface, "down")
		_, err = command.Output()
		if err != nil {
			if err.Error() == "exit status 255" {
				fmt.Println("Warning : run the tool using sudo")
				os.Exit(1)
			}
			fmt.Println("Err: ", err.Error())
			os.Exit(1)
		}

		time.Sleep(10 * time.Second) // wait for interface down event
		command = exec.Command("ifconfig", iface, "hw", "ether", mac)
		_, err = command.CombinedOutput()
		if err != nil {
			fmt.Println("Err: ", err.Error())
			os.Exit(1)
		}
		command = exec.Command("ifconfig", iface, "up")
		_, err = command.CombinedOutput()
		if err != nil {
			fmt.Println("Err: ", err.Error())
			os.Exit(1)
		}
		fmt.Println(fmt.Sprintf("New MAC  Address  =>  %s", mac))

	},
}

func init() {
	chageMacCmd.Flags().StringVarP(&iface, "interface", "i", "", "Inteface name")
	chageMacCmd.Flags().StringVarP(&mac, "new-mac", "m", "", "New mac address ")
	if err := chageMacCmd.MarkFlagRequired("interface"); err != nil {
		fmt.Println(err.Error(), " \nuse net change-mac -h  for more detail")
		os.Exit(1)
	}
}
