package go_crypto_otc

import (
	"time"
)





type DepthRecord struct {
	Price,
	Amount float64
}

type DepthRecords []DepthRecord

func (dr DepthRecords) Len() int {
	return len(dr)
}

func (dr DepthRecords) Swap(i, j int) {
	dr[i], dr[j] = dr[j], dr[i]
}

func (dr DepthRecords) Less(i, j int) bool {
	return dr[i].Price < dr[j].Price
}

type Depth struct {
	Currency         Currency
	UTime        time.Time
	AskList,
	BidList DepthRecords
}


func (self *Depth) AggDep() *Depth{


	ask_depth := make(DepthRecords,0)

	ask_depth=append(ask_depth, self.AskList[0])
	for i := 1; i < len(self.AskList); i++ {
		if self.AskList[i].Price==ask_depth[len(ask_depth)-1].Price {
			ask_depth[len(ask_depth)-1].Amount += self.AskList[i].Amount
		}else{
			ask_depth=append(ask_depth, self.AskList[i])
		}
	}

	bid_depth := make(DepthRecords,0)

	bid_depth=append(bid_depth, self.BidList[0])
	for i := 1; i < len(self.BidList); i++ {
		if self.BidList[i].Price==bid_depth[len(bid_depth)-1].Price {
			bid_depth[len(bid_depth)-1].Amount += self.BidList[i].Amount
		}else{
			bid_depth=append(bid_depth, self.BidList[i])
		}
	}

	dep:=&Depth{
		self.Currency,
		self.UTime,
		ask_depth,
		bid_depth,
	}
	return dep
}