package go_crypto_otc

import "strings"


// ETH_BTC --> ethbtc
type Symbols map[CurrencyPair]string

// huobi.com --> symbols
type ExSymbols map[string]Symbols

var exSymbols ExSymbols

func GetExSymbols(exName string) Symbols {
	ret, ok := exSymbols[exName]
	if !ok {
		return nil
	}
	return ret
}

func RegisterExSymbol(exName string, pair CurrencyPair) {
	if exSymbols == nil {
		exSymbols = make(ExSymbols)
	}

	if _, ok := exSymbols[exName]; !ok {
		exSymbols[exName] = make(Symbols)
	}

	exSymbols[exName][pair] = pair.ToSymbol("")
}

type Currency struct {
	Symbol string
	Desc   string
}

func (c Currency) String() string {
	return c.Symbol
}

// A->B(A兑换为B)
type CurrencyPair struct {
	CurrencyA Currency
	CurrencyB Currency
}

var (

	CNY     = Currency{"CNY", ""}
	USD     = Currency{"USD", ""}
	USDT    = Currency{"USDT", ""}

	BTC     = Currency{"BTC", ""}
	LTC     = Currency{"LTC", ""}
	ETH     = Currency{"ETH", ""}


	//currency pair

	BTC_CNY  = CurrencyPair{BTC, CNY}
	LTC_CNY  = CurrencyPair{LTC, CNY}
	ETH_CNY  = CurrencyPair{ETH, CNY}




	BTC_USD = CurrencyPair{BTC, USD}
	LTC_USD = CurrencyPair{LTC, USD}
	ETH_USD = CurrencyPair{ETH, USD}

	BTC_USDT = CurrencyPair{BTC, USDT}
	LTC_USDT = CurrencyPair{LTC, USDT}
	ETH_USDT = CurrencyPair{ETH, USDT}


	LTC_BTC = CurrencyPair{LTC, BTC}
	ETH_BTC = CurrencyPair{ETH, BTC}

)

func (c CurrencyPair) String() string {
	return c.ToSymbol("_")
}


func NewCurrency(symbol, desc string) Currency {
	switch symbol {
	case "cny", "CNY":
		return CNY
	case "usdt", "USDT":
		return USDT
	case "usd", "USD":
		return USD
	default:
		return Currency{strings.ToUpper(symbol), desc}
	}
}

func NewCurrencyPair(currencyA Currency, currencyB Currency) CurrencyPair {
	return CurrencyPair{currencyA, currencyB}
}

func NewCurrencyPair2(currencyPairSymbol string) CurrencyPair {
	currencys := strings.Split(currencyPairSymbol, "_")
	if len(currencys) == 2 {
		return CurrencyPair{NewCurrency(currencys[0], ""),
			NewCurrency(currencys[1], "")}
	}
	panic("symbol error!!!!")
}

func (pair CurrencyPair) ToSymbol(joinChar string) string {
	return strings.Join([]string{pair.CurrencyA.Symbol, pair.CurrencyB.Symbol}, joinChar)
}

func (pair CurrencyPair) ToSymbol2(joinChar string) string {
	return strings.Join([]string{pair.CurrencyB.Symbol, pair.CurrencyA.Symbol}, joinChar)
}

func (pair CurrencyPair) AdaptUsdtToUsd() CurrencyPair {
	CurrencyB := pair.CurrencyB
	if pair.CurrencyB == USDT {
		CurrencyB = USD
	}
	return CurrencyPair{pair.CurrencyA, CurrencyB}
}

func (pair CurrencyPair) AdaptUsdToUsdt() CurrencyPair {
	CurrencyB := pair.CurrencyB
	if pair.CurrencyB == USD {
		CurrencyB = USDT
	}
	return CurrencyPair{pair.CurrencyA, CurrencyB}
}


//for to symbol lower , Not practical '==' operation method
func (pair CurrencyPair) ToLower() CurrencyPair {
	return CurrencyPair{NewCurrency(strings.ToLower(pair.CurrencyA.Symbol), ""),
		NewCurrency(strings.ToLower(pair.CurrencyB.Symbol), "")}
}

type CurrencyPair2 struct {
	CurrencyPair
}

func (pair CurrencyPair) Reverse() CurrencyPair2 {
	return CurrencyPair2{CurrencyPair{pair.CurrencyB, pair.CurrencyA}}
}

func (pair CurrencyPair2) ToSymbol(joinChar string) string {
	return strings.Join([]string{pair.CurrencyA.Symbol, pair.CurrencyB.Symbol}, joinChar)
}
