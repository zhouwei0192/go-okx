package market

import "github.com/zhouwei0192/go-okx/rest/api"

func NewGetMarkPriceCandles(param *GetCandlesParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/market/mark-price-candles",
		Method: api.MethodGet,
		Param:  param,
	}, &GetCandlesResponse{}
}
