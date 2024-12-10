package main

import (
	"log"

	"github.com/zhouwei0192/go-okx/examples/rest"
	"github.com/zhouwei0192/go-okx/rest/api/trade"
)

func main() {
	param := &trade.PostAmendOrderParam{
		InstId: "OKB-USDT",
		OrdId:  "515102546340442112",
		NewSz:  "1.5",
	}
	req, resp := trade.NewPostAmendOrder(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.PostAmendOrderResponse))
}
