package main

import (
	"log"

	"github.com/zhouwei0192/go-okx/examples/rest"
	"github.com/zhouwei0192/go-okx/rest/api/market"
)

func main() {
	param := &market.GetTickerParam{
		InstId: "ETH-USDT",
	}
	req, resp := market.NewGetTicker(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*market.GetTickerResponse))
}
