package main

import (
	"log"

	"github.com/zhouwei0192/go-okx/examples/rest"
	"github.com/zhouwei0192/go-okx/rest/api/account"
)

func main() {
	param := &account.GetBalanceParam{}
	req, resp := account.NewGetBalance(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.GetBalanceResponse))
}
