package zb

import (
	"github.com/Vniex/tkTools"
	"go-crypto-otc"
	"testing"
)

var httpClient=tkTools.NewHttpClient(3,"http://127.0.0.1:1087")
var ex=NewZB(httpClient)

func TestZB_GetExchangeName(t *testing.T) {
	name:=ex.GetExchangeName()
	t.Log(name)
}

func TestZB_GetDepth(t *testing.T) {
	dep,err:=ex.GetDepth(10,go_crypto_otc.QC)
	if err!=nil{
		t.Log(err)
	}else{
		t.Log(dep.AskList)
		t.Log(dep.BidList)
	}



}
