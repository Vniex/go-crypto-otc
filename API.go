package go_crypto_otc

// api interface

type API interface {

	GetDepth(size int, currency Currency) (*Depth, error)

	GetExchangeName() string

	GetWithdrawalsFee(currency Currency)float64
}
