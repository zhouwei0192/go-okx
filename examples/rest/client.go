package rest

import (
	"github.com/zhouwei0192/go-okx/examples"
	rc "github.com/zhouwei0192/go-okx/rest"
)

// 敏感信息申请的模拟盘KEY，不确定何时会删除
var TestClient = rc.New("", examples.Auth, nil)
