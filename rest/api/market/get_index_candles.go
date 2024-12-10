package market

import "github.com/zhouwei0192/go-okx/rest/api"

func NewGetIndexCandles(param *GetCandlesParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/market/index-candles",
		Method: api.MethodGet,
		Param:  param,
	}, &GetCandlesResponse{}
}
