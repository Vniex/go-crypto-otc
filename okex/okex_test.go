package okex

import (
	"go-crypto-otc"
	"github.com/Vniex/tkTools"
	"testing"
)

var httpClient=tkTools.NewHttpClient(3,"http://127.0.0.1:1087")
var ex=NewOKEX(httpClient)

func TestOKEX_GetExchangeName(t *testing.T) {
	name:=ex.GetExchangeName()
	t.Log(name)
}

func TestOKEX_GetDepth(t *testing.T) {
	dep,err:=ex.GetDepth(4,go_crypto_otc.USDT)
	t.Log(err)
	t.Log(dep.AskList)
	t.Log(dep.BidList)

}

