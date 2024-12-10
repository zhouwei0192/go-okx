package main

import (
	"log"

	"github.com/zhouwei0192/go-okx/examples/rest"
	"github.com/zhouwei0192/go-okx/rest/api/market"
)

func main() {
	param := &market.GetIndexComponentsParam{
		Index: "BTC-USDT",
	}
	req, resp := market.NewGetIndexComponents(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*market.GetIndexComponentsResponse))
}
