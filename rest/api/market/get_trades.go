package market

import "github.com/zhouwei0192/go-okx/rest/api"

func NewGetTrades(param *GetTradesParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/market/trades",
		Method: api.MethodGet,
		Param:  param,
	}, &GetTradesResponse{}
}

type GetTradesParam struct {
	InstId string `url:"instId"`
	Limit  int    `url:"limit,omitempty"`
}

type GetTradesResponse struct {
	api.Response
	Data []Trade `json:"data"`
}

type Trade struct {
	InstId  string `json:"instId"`
	TradeId string `json:"tradeId"`
	Px      string `json:"px"`
	Sz      string `json:"sz"`
	Side    string `json:"side"`
	Ts      int64  `json:"ts,string"`
}
