package comm

type JsonObj map[string]interface{}

type ExchangeType string

const (
	Huobi ExchangeType = "huobi"
)

var (
	Consts        = []string{"BTC", "LTC", "M", "M5", "M15", "M30", "H", "D", "W"}
	ExchangeTypes = []ExchangeType{Huobi /*OkCoinCn, Poloniex, Btcc, Chbtc, OkcoinFuture*/}
)
