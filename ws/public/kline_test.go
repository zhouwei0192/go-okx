package public

import (
	"github.com/zhouwei0192/go-okx/ws"
	"log"
	"testing"
)

func TestSubscribeKline(t *testing.T) {
	args := make([]ws.Args, 0)
	args = append(args, ws.Args{
		Channel: "candle1m",
		InstId:  "BTC-USDT",
	})
	handler := func(c EventKline) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := SubscribeKline(args, handler, handlerError, false); err != nil {
		panic(err)
	}
	select {}
}
