package zb
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
	ZB_BTC="btc_cny"
	ZB_USDT="usdt_cny"
	ZB_QC="qc_cny"
	OTC_URL="https://vip.zb.cn/api/web/otc/V1_0_0/getOnlineAdList?marketName=%s&type=%s&priceSort=%s&pageIndex=1&pageSize=10"
)

type ZB struct {
	httpClient *http.Client
}

func NewZB(httpClient *http.Client)*ZB{
	return &ZB{
		httpClient:httpClient,
	}
}

func (self *ZB)GetExchangeName() string{
	return "zb.com"
}

func (self *ZB)GetDepth(size int, currency Currency) (*Depth,error){

	currencyID:=self.currency2Id(currency)
	bid_URL:=fmt.Sprintf(OTC_URL,currencyID,"1","1")
	ask_URL:=fmt.Sprintf(OTC_URL,currencyID,"2","2")
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
	return dep.AggDep(),nil
}
func (self  *ZB)ParseDepth(depmap map[string]interface{}) (DepthRecords,error){
	resMsgMap:=depmap["resMsg"].(map[string]interface{})
	if resMsgMap["code"].(float64)!=1000{
		return nil,errors.New(resMsgMap["message"].(string))
	}
	dataAllmap := depmap["datas"].(map[string]interface{})
	datamap:=dataAllmap["list"].([]interface{})
	depth := make(DepthRecords,0)
	for _, r := range datamap {
		var dr DepthRecord
		rr := r.(map[string]interface{})
		dr.Price = tkTools.ToFloat64(rr["price"])
		dr.Amount = tkTools.ToFloat64(rr["remainAmount"])
		depth = append(depth, dr)
	}
	return depth,nil
}

func (self *ZB)currency2Id(currency Currency)string{
	switch currency.Symbol {
	case "BTC":
		return ZB_BTC
	case "USDT":
		return ZB_USDT
	case "QC":
		return ZB_QC


	}
	panic("currency error")
}


func (self *ZB)GetWithdrawalsFee(currency Currency)float64{
	switch currency {
	case USDT:
		return 5
	case BTC:
		return 0.001
	}
	panic("currency error")
}