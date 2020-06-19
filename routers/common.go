package routers

const (
	identityKey string = "guid"

	// login type
	GuestLogin = "guest"
	PhoneLogin = "phone"

	// role
	UserAdmin string = "admin"
)

// exchange types

type ExchangeType string

const (
	Huobi ExchangeType = "huobi"
	// OkCoinCn     ExchangeType = "okcoin.cn"
	// Poloniex     ExchangeType = "poloniex"
	// Btcc         ExchangeType = "btcc"
	// Chbtc        ExchangeType = "chbtc"
	// OkcoinFuture ExchangeType = "okcoin.future"
	// OandaV20     ExchangeType = "oanda.v20"
)

var (
	Consts        = []string{"BTC", "LTC", "M", "M5", "M15", "M30", "H", "D", "W"}
	ExchangeTypes = []ExchangeType{Huobi /*OkCoinCn, Poloniex, Btcc, Chbtc, OkcoinFuture*/}
)