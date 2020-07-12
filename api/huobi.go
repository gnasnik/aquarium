package api

import (
	"time"

	"github.com/frankffenn/aquarium/utils/log"
	"github.com/google/uuid"
	cli "github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/getrequest"
	"github.com/huobirdcenter/huobi_golang/pkg/postrequest"
	"golang.org/x/xerrors"
)

type Huobi struct {
	options      Options
	symbolMap    map[string]string
	minAmountMap map[string]float64

	common  *cli.CommonClient
	account *cli.AccountClient
	market  *cli.MarketClient
	order   *cli.OrderClient
	wallet  *cli.WalletClient

	lastSleep int64
	lastTimes int64
}

func NewHuobi(opts ...Option) *Huobi {
	options := newOption()

	for _, o := range opts {
		o(&options)
	}

	return &Huobi{
		options: options,
		symbolMap: map[string]string{
			"btcusdt": "1",
		},
		minAmountMap: map[string]float64{
			"btcusdt": 0.001,
		},
		common:    new(cli.CommonClient).Init(options.Host),
		market:    new(cli.MarketClient).Init(options.Host),
		account:   new(cli.AccountClient).Init(options.AccessKey, options.SecretKey, options.Host),
		order:     new(cli.OrderClient).Init(options.AccessKey, options.SecretKey, options.Host),
		wallet:    new(cli.WalletClient).Init(options.AccessKey, options.SecretKey, options.Host),
		lastSleep: time.Now().UnixNano(),
	}
}

func (h *Huobi) GetType() string {
	return h.options.Type
}

func (h *Huobi) GetName() string {
	return h.options.Name
}

func (h *Huobi) AutoSleep() {
	now := time.Now().UnixNano()
	limit := h.options.Limit

	interval := int64(1e+9/limit*h.lastTimes - (now - h.lastSleep))
	if interval > 0 {
		time.Sleep(time.Duration(interval))
	}

	h.lastTimes = 0
	h.lastSleep = now
}

func (h *Huobi) GetMinAmount(stock string) float64 {
	return h.minAmountMap[stock]
}

func (h *Huobi) GetAccountInfo() (interface{}, error) {
	rsp, err := h.account.GetAccountInfo()
	if err != nil {
		log.Err("get account info failed, %v", err)
		return nil, err
	}
	return rsp, nil
}

func (h *Huobi) GetAccountBalance(id string) (interface{}, error) {
	rsp, err := h.account.GetAccountBalance(id)
	if err != nil {
		log.Err("get account balance faield, %v", err)
		return nil, err
	}
	return rsp, nil
}

func (h *Huobi) PlaceOrder(accountID, symbol, orderType, amount, price string) (interface{}, error) {
	req := &postrequest.PlaceOrderRequest{
		AccountId:     accountID,
		Symbol:        symbol,
		Type:          orderType,
		Amount:        amount,
		Price:         price,
		Source:        "spot-api",
		ClientOrderId: uuid.New().String(),
		// StopPrice:     "",
		// Operator:      "",
	}
	rsp, err := h.order.PlaceOrder(req)
	if err != nil {
		log.Err("place order failed,%v", err)
		return nil, err
	}

	return rsp, nil
}

func (h *Huobi) GetOrder(orderID string) (interface{}, error) {
	rsp, err := h.order.GetOrderById(orderID)
	if err != nil {
		log.Err("get order failed, %v", orderID)
		return nil, err
	}
	return rsp.Data, nil
}

func (h *Huobi) GetOrders(symbol string) (interface{}, error) {
	if _, ok := h.symbolMap[symbol]; !ok {
		log.Err("GetOrders() error, unrecognized symbol: ", symbol)
		return nil, xerrors.New("unrecognized symbol")
	}
	req := new(getrequest.GetRequest).Init()
	req.AddParam("symbol", symbol)
	req.AddParam("states", "submitted,partial-filled")
	rsp, err := h.order.GetHistoryOrders(req)
	if err != nil {
		return nil, err
	}

	return rsp.Data, nil
}

func (h *Huobi) GetTrades(symbol string) (interface{}, error) {
	if _, ok := h.symbolMap[symbol]; !ok {
		log.Err("GetTrades() error, unrecognized symbol: ", symbol)
		return nil, xerrors.New("unrecognized symbol")
	}
	req := new(getrequest.GetRequest).Init()
	req.AddParam("symbol", symbol)
	req.AddParam("states", "filled")
	rsp, err := h.order.GetHistoryOrders(req)
	if err != nil {
		return nil, err
	}

	return rsp.Data, nil
}

func (h *Huobi) CancelOrder(orderID string) error {
	rsp, err := h.order.CancelOrderById(orderID)
	if err != nil {
		log.Err("cancel order failed,%v", rsp.ErrorMessage)
		return err
	}
	return nil
}

func (h *Huobi) GetDepth(symbol, step string, opts ...CallOption) (interface{}, error) {
	callOpts := newCallOption()

	for _, o := range opts {
		o(&callOpts)
	}

	reqOpt := getrequest.GetDepthOptionalRequest{
		Size: callOpts.Size,
	}

	rsp, err := h.market.GetDepth(symbol, step, reqOpt)
	if err != nil {
		log.Err("cancel order failed %v", err)
		return nil, err
	}

	return rsp, nil
}

func (h *Huobi) GetCandlestick(symbol, period string, opts ...CallOption) (interface{}, error) {
	callOpts := newCallOption()

	for _, o := range opts {
		o(&callOpts)
	}

	reqOpt := getrequest.GetCandlestickOptionalRequest{
		Size:   callOpts.Size,
		Period: period,
	}

	rsp, err := h.market.GetCandlestick(symbol, reqOpt)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

var _ Exchange = &Huobi{}
