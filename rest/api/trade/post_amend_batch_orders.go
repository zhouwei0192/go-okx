package trade

import "github.com/zhouwei0192/go-okx/rest/api"

func NewPostAmendBatchOrders(param []*PostAmendOrderParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/amend-batch-orders",
		Method: api.MethodPost,
		Param:  param,
	}, &PostAmendOrderResponse{}
}
