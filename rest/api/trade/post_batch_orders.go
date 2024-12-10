package trade

import "github.com/zhouwei0192/go-okx/rest/api"

func NewPostBatchOrders(param []*PostOrderParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/batch-orders",
		Method: api.MethodPost,
		Param:  param,
	}, &PostOrderResponse{}
}
