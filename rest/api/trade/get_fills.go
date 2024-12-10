package trade

import (
	"github.com/zhouwei0192/go-okx/rest/api"
)

func NewGetFills(param *GetFillsParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/fills",
		Method: api.MethodGet,
		Param:  param,
	}, &GetFillsResponse{}
}

type GetFillsParam struct {
	api.PagingParam
	InstType string `url:"instType,omitempty"`
	Uly      string `url:"uly,omitempty"`
	InstId   string `url:"instId,omitempty"`
	OrdId    string `url:"ordId,omitempty"`
}

type GetFillsResponse struct {
	api.Response
	Data []Fill `json:"data"`
}

type Fill struct {
	InstType string `json:"instType"`
	InstId   string `json:"instId"`
	TradeId  string `json:"tradeId"`
	OrdId    string `json:"ordId"`
	ClOrdId  string `json:"clOrdId"`
	BillId   string `json:"billId"`
	Tag      string `json:"tag"`
	FillPx   string `json:"fillPx"`
	FillSz   string `json:"fillSz"`
	Side     string `json:"side"`
	PosSide  string `json:"posSide"`
	ExecType string `json:"execType"`
	FeeCcy   string `json:"feeCcy"`
	Fee      string `json:"fee"`
	Ts       int64  `json:"ts,string"`
}
