package main

import (
	"fmt"
	"github.com/depocket/lab/apps"
)

func main() {
	service, err := apps.ServiceRegisters.New("bsc_lido")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(service.Code())
	fmt.Println(service.FetchBalances())
}
