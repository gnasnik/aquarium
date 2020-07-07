package api

import (
	modacc "github.com/huobirdcenter/huobi_golang/pkg/response/account"
	modord "github.com/huobirdcenter/huobi_golang/pkg/response/order"
)

// Exchange interface
type Exchange interface {
	// Log(...interface{})
	GetType() string
	GetName() string
	// SetLimit(times interface{}) float64
	AutoSleep()
	GetMinAmount(stock string) float64
	GetAccountBalance(id string) *modacc.AccountBalance
	PlaceOrder(accountID, symbol, orderType, amount, price string) *modord.PlaceOrderResponse
	GetOrder(stockType, id string) interface{}
	GetOrders(stockType string) interface{}
	GetTrades(stockType string) interface{}
	// CancelOrder(order Order) bool
	GetTicker(stockType string, sizes ...interface{}) interface{}
	GetRecords(stockType, period string, sizes ...interface{}) interface{}
}

var (
	constructor = map[string]func(Option) Exchange{}
)
