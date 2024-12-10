package market

import "github.com/zhouwei0192/go-okx/rest/api"

func NewGetBooks(param *GetBooksParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/market/books",
		Method: api.MethodGet,
		Param:  param,
	}, &GetBooksResponse{}
}

type GetBooksParam struct {
	InstId string `url:"instId"`
	Sz     string `url:"sz,omitempty"`
}

type GetBooksResponse struct {
	api.Response
	Data []Book `json:"data"`
}

type Book struct {
	Asks [][]string `json:"asks"`
	Bids [][]string `json:"bids"`
	Ts   int64      `json:"ts,string"`
}
