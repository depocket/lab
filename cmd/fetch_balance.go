package main

import (
	"fmt"

	"github.com/depocket/go-client"
	_ "github.com/depocket/lab/apps"
	"github.com/depocket/lab/constants"
	"github.com/depocket/lab/core"
)

func main() {
	baseUrl := "https://sdk-stg.depocket.io/api/v1/"
	cl := client.NewClient(nil, &baseUrl)

	// Binance
	bscApps := []string{"pancakeswap"}
	for _, name := range bscApps {
		service, err := core.ServiceRegisters.New(name)
		if err != nil {
			fmt.Println(err)
		}
		err = service.InjectDependencies(constants.ChainBinance, cl)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(service.ApplicationManifest())
		fmt.Println(service.Balances(constants.ChainBinance, name, "0x0000"))
	}

	// Ethereum
	ethApps := []string{"curve"}
	for _, name := range ethApps {
		service, err := core.ServiceRegisters.New(name)
		if err != nil {
			fmt.Println(err)
		}
		err = service.InjectDependencies(constants.ChainEthereum, cl)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(service.ApplicationManifest())
		fmt.Println(service.Balances(constants.ChainEthereum, name, "0x7a16ff8270133f063aab6c9977183d9e72835428"))
	}
}
