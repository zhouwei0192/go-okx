package account

import "github.com/zhouwei0192/go-okx/rest/api"

func NewGetBillsArchive(param *GetBillsParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/account/bills-archive",
		Method: api.MethodGet,
		Param:  param,
	}, &GetBillsResponse{}
}
