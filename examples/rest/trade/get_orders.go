package main

import (
	"log"

	"github.com/zhouwei0192/go-okx/examples/rest"
	"github.com/zhouwei0192/go-okx/rest/api/trade"
)

func main() {
	param := &trade.GetOrderParam{
		InstId: "OKB-USDT",
		OrdId:  "501163171776954368",
	}
	req, resp := trade.NewGetOrder(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.GetOrderResponse))
}
