package net

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

type portScanner struct {
	address         string
	numberOfWorkers int
}

func (p *portScanner) workers(ports chan int, wg *sync.WaitGroup) {

	for port := range ports {
		add := fmt.Sprintf("%s:%d", p.address, port)
		conn, err := net.DialTimeout("tcp", add, 400*time.Millisecond)
		if err != nil {
			wg.Done()
			continue
		}
		fmt.Printf("port %d is open \n", port)
		conn.Close()
		wg.Done()

	}
}

var (
	addr     string
	nworkers int
)

func scan(address string, numberOfWorkers int) {
	scanner := portScanner{
		address:         address,
		numberOfWorkers: numberOfWorkers,
	}
	ports := make(chan int, numberOfWorkers)
	var wg sync.WaitGroup
	for i := 0; i < cap(ports); i++ {
		go scanner.workers(ports, &wg)
	}

	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)
}

var scanCmd = &cobra.Command{
	Use:   "port-scan",
	Short: "port-scan used to scan open ports ",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		scan(addr, nworkers)
	},
}

func init() {

	scanCmd.Flags().StringVarP(&addr, "address", "a", "", "target ip address")
	scanCmd.Flags().IntVarP(&nworkers, "workers", "w", 100, "number of thread you want to run  the more number you give the more faster result default 100")

	if err := scanCmd.MarkFlagRequired("address"); err != nil {
		fmt.Println(err.Error(), " use net ping -h  for more detail")
		os.Exit(1)
	}
}
