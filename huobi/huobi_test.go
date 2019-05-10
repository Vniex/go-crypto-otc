package huobi

import (
	"github.com/Vniex/tkTools"
	"testing"
	"go-crypto-otc"
)

var httpClient=tkTools.NewHttpClient(3,"http://127.0.0.1:1087")
var huobi=NewHuobi(httpClient)

func TestHuobi_GetExchangeName(t *testing.T) {
	name:=huobi.GetExchangeName()
	t.Log(name)
}

func TestHuobi_GetDepth(t *testing.T) {
	dep,err:=huobi.GetDepth(50,go_crypto_otc.USDT)
	t.Log(err)
	t.Log(dep.AskList)
	t.Log(dep.BidList)

}