package trade

import (
	"github.com/zhouwei0192/go-okx/rest/api"
)

func NewGetOrdersHistory(param *GetOrdersQueryParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/orders-history",
		Method: api.MethodGet,
		Param:  param,
	}, &GetOrderResponse{}
}
