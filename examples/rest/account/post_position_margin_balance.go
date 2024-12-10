package main

import (
	"log"

	"github.com/zhouwei0192/go-okx/examples/rest"
	"github.com/zhouwei0192/go-okx/rest/api"
	"github.com/zhouwei0192/go-okx/rest/api/account"
)

func main() {
	param := &account.PostPositionMarginBalanceParam{
		InstId:  "BTC-USDT-SWAP",
		PosSide: api.PosSideNet,
		Type:    api.TypeReduce,
		Amt:     "1",
	}
	req, resp := account.NewPostPositionMarginBalance(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.PostPositionMarginBalanceResponse))
}
