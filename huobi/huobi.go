package huobi

import (
	"net/http"
	. "go-crypto-otc"
	"github.com/Vniex/tkTools/httpUtils"
	"fmt"
	"time"
	"errors"
	"github.com/Vniex/tkTools"
)


const(
	Huobi_BTC="1"
	Huobi_USDT="2"
	OTC_URL="https://otc-api.hbg.com/v1/data/trade-market?country=37&currency=1&payMethod=0&currPage=1&coinId=%s&tradeType=%s&blockType=general&online=1"
)

type Huobi struct {
	httpClient *http.Client
}

func NewHuobi(httpClient *http.Client)*Huobi{
	return &Huobi{
		httpClient:httpClient,
	}
}

func (self *Huobi)GetExchangeName() string{
	return "huobi.com"
}

func (self *Huobi)GetDepth(size int, currency Currency) (*Depth,error){

	currencyID:=self.currency2Id(currency)
	bid_URL:=fmt.Sprintf(OTC_URL,currencyID,"buy")
	ask_URL:=fmt.Sprintf(OTC_URL,currencyID,"sell")
	bid_map,err:=httpUtils.HttpGet(self.httpClient,bid_URL)
	if err!=nil{
		return nil,err
	}
	ask_map,err:=httpUtils.HttpGet(self.httpClient,ask_URL)
	if err!=nil{
		return nil,err
	}
	ask_dep,err:=self.ParseDepth(ask_map)
	if err!=nil{
		return nil,err
	}
	bid_dep,err:=self.ParseDepth(bid_map)
	if err!=nil{
		return nil,err
	}
	dep:=&Depth{
		currency,
		time.Now(),
		ask_dep,
		bid_dep,
	}
	return dep,nil
}
func (self  *Huobi)ParseDepth(depmap map[string]interface{}) (DepthRecords,error){
	if depmap["code"].(float64)!=200{
		return nil,errors.New(depmap["message"].(string))
	}
	datamap := depmap["data"].([]interface{})
	depth := make(DepthRecords,0)
	for _, r := range datamap {
		var dr DepthRecord
		rr := r.(map[string]interface{})
		dr.Price = tkTools.ToFloat64(rr["price"])
		dr.Amount = tkTools.ToFloat64(rr["tradeCount"])
		depth = append(depth, dr)
	}
	return depth,nil
}

func (self *Huobi)currency2Id(currency Currency)string{
	switch currency.Symbol {
	case "BTC":
		return Huobi_BTC
	case "USDT":
		return Huobi_USDT

	}
	panic("currency error")
}


func (self *Huobi)GetWithdrawalsFee(currency Currency)float64{
	switch currency {
	case USDT:
		return 5
	case BTC:
		return 0.0005
	}
	panic("currency error")
}