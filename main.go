package main

import (
	"context"
	"fmt"
	"github.com/depocket/lab/sdk"
)

func main() {
	//service, err := apps.ServiceRegisters.New("ethereum_lido")
	//if err != nil {
	//	fmt.Println(err)
	//}
	baseUrl := "https://sdk-stg.depocket.io/api/v1/"
	client := sdk.NewClient(nil, &baseUrl)
	ethTokens, _, err := client.Pools.ListByProjectCode(context.TODO(), "bsc_pancakeswap", &sdk.PoolListOptions{
		ListOptions: sdk.ListOptions{
			Chain: "bsc",
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ethTokens)
	//err = service.BeforeFetchData(client)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(service.Code())
	//fmt.Println(service.FetchBalances())
}
