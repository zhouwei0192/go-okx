package main

import (
	"log"

	"github.com/zhouwei0192/go-okx/examples/rest"
	"github.com/zhouwei0192/go-okx/rest/api"
	"github.com/zhouwei0192/go-okx/rest/api/trade"
)

func main() {
	param := &trade.PostOrderParam{
		InstId:  "OKB-USDT",
		TdMode:  api.TdModeCash,
		Side:    api.SideBuy,
		OrdType: api.OrdTypeLimit,
		Sz:      "9",
		Px:      "5",
	}
	req, resp := trade.NewPostOrder(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.PostOrderResponse))
}
