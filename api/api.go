package api

// Exchange interface
type Exchange interface {
	// Log(...interface{})
	GetType() string
	GetName() string
	AutoSleep()
	GetMinAmount(stock string) float64
	GetAccountBalance(id string) (interface{}, error)
	PlaceOrder(accountID, symbol, orderType, amount, price string) (interface{}, error)
	GetOrder(orderID string) (interface{}, error)
	GetOrders(stockType string) (interface{}, error)
	GetTrades(stockType string) (interface{}, error)
	CancelOrder(orderID string) error
	GetDepth(symbol, step string, opts ...CallOption) (interface{}, error)
	GetCandlestick(symbol, period string, opts ...CallOption) (interface{}, error)
}

var (
	constructor = map[string]func(Option) Exchange{}
)
