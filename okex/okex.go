package okex
import (
	"net/http"
	. "go-crypto-otc"
	"github.com/Vniex/tkTools/httpUtils"
	"fmt"
	"time"
	"errors"
	"github.com/Vniex/tkTools"
	"sort"
)


const(
	OKEX_BTC="btc"
	OKEX_USDT="usdt"
	OTC_URL="https://www.okex.com/v3/c2c/tradingOrders/book?baseCurrency=%s&quoteCurrency=cny&userType=certified&paymentMethod=all"
)

type OKEX struct {
	httpClient *http.Client
}

func NewOKEX(httpClient *http.Client)*OKEX{
	return &OKEX{
		httpClient:httpClient,
	}
}

func (self *OKEX)GetExchangeName() string{
	return "okex.com"
}

func (self *OKEX)GetDepth(size int, currency Currency) (*Depth,error){

	currencyID:=self.currency2Id(currency)

	dep_URL:=fmt.Sprintf(OTC_URL,currencyID)
	dep_map,err:=httpUtils.HttpGet(self.httpClient,dep_URL)
	if err!=nil{
		return nil,err
	}
	dep,err:=self.ParseFullDepth(dep_map,currency)
	if err!=nil{
		return nil,err
	}
	sort.Sort(DepthRecords(dep.AskList))
	sort.Sort(sort.Reverse(DepthRecords(dep.BidList)))

	return dep,nil
}

func (self *OKEX)ParseFullDepth(depmap map[string]interface{},currency Currency)(*Depth,error){
	if depmap["code"].(float64)!=0{
		return nil,errors.New(depmap["detailMsg"].(string))
	}
	datamap := depmap["data"].(map[string]interface{})
	dep:=&Depth{
		currency,
		time.Now(),
		self.ParseDepth(datamap["sell"].([]interface{})),
		self.ParseDepth(datamap["buy"].([]interface{})),
	}
	return dep.AggDep(),nil

}

func (self  *OKEX)ParseDepth(deplist []interface{}) (DepthRecords){

	depth := make(DepthRecords,0)
	for _, r := range deplist {
		var dr DepthRecord
		rr := r.(map[string]interface{})
		dr.Price = tkTools.ToFloat64(rr["price"])
		dr.Amount = tkTools.ToFloat64(rr["availableAmount"])
		depth = append(depth, dr)
	}
	return depth
}

func (self *OKEX)currency2Id(currency Currency)string{
	switch currency.Symbol {
	case "BTC":
		return OKEX_BTC
	case "USDT":
		return OKEX_USDT

	}
	panic("currency error")
}

func (self *OKEX)GetWithdrawalsFee(currency Currency)float64{
	switch currency {
	case USDT:
		return 2
	case BTC:
		return 0.0005
	}
	panic("currency error")
}