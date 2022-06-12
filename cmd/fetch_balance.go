package main

import (
	"fmt"
	"github.com/depocket/go-client"
	_ "github.com/depocket/lab/apps"
	"github.com/depocket/lab/constants"
	"github.com/depocket/lab/core"
)

func main() {
	service, err := core.ServiceRegisters.New("pancakeswap")
	if err != nil {
		fmt.Println(err)
	}
	baseUrl := "https://sdk-stg.depocket.io/api/v1/"
	cl := client.NewClient(nil, &baseUrl)
	err = service.InjectDependencies(constants.ChainBinance, cl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(service.ApplicationManifest())
	fmt.Println(service.Balances(constants.ChainBinance, "pancakeswap", "0x0000"))
}
