package main

import (
	"log"

	"github.com/zhouwei0192/go-okx/examples/rest"
	"github.com/zhouwei0192/go-okx/rest/api/asset"
)

func main() {
	param := &asset.GetDepositAddressParam{
		Ccy: "BTC",
	}
	req, resp := asset.NewGetDepositAddress(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*asset.GetDepositAddressResponse))
}
