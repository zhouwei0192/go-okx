package main

import (
	"log"

	"github.com/zhouwei0192/go-okx/examples/rest"
	"github.com/zhouwei0192/go-okx/rest/api"
	"github.com/zhouwei0192/go-okx/rest/api/trade"
)

func main() {
	param := &trade.PostClosePositionParam{
		InstId:  "OKB-USDT",
		MgnMode: api.MgnModeCross,
		Ccy:     "OKB",
	}
	req, resp := trade.NewPostClosePosition(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.PostClosePositionResponse))
}
