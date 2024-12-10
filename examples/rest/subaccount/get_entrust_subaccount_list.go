package main

import (
	"log"

	"github.com/zhouwei0192/go-okx/examples/rest"
	"github.com/zhouwei0192/go-okx/rest/api/subaccount"
)

func main() {
	param := &subaccount.GetEntrustSubaccountListParam{}
	req, resp := subaccount.NewGetEntrustSubaccountList(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*subaccount.GetEntrustSubaccountListResponse))
}
