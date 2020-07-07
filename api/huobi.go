package api

import (
	"time"

	"github.com/frankffenn/aquarium/utils/log"
	"github.com/google/uuid"
	conf "github.com/huobirdcenter/huobi_golang/config"
	cli "github.com/huobirdcenter/huobi_golang/pkg/client"
	modpr "github.com/huobirdcenter/huobi_golang/pkg/postrequest"
	modacc "github.com/huobirdcenter/huobi_golang/pkg/response/account"
	modord "github.com/huobirdcenter/huobi_golang/pkg/response/order"
)

type Huobi struct {
	opts         Options
	minAmountMap map[string]float64

	common  *cli.CommonClient
	account *cli.AccountClient
	market  *cli.MarketClient
	order   *cli.OrderClient
	wallet  *cli.WalletClient

	lastSleep int64
	lastTimes int64
}

func NewHuobi(opt ...Option) *Huobi {
	opts := newOption()
	if opts.Host == "" {
		opts.Host = conf.Host
	}

	return &Huobi{
		opts: opts,
		minAmountMap: map[string]float64{
			"BTC/CNY": 0.001,
		},
		common:    new(cli.CommonClient).Init(opts.Host),
		market:    new(cli.MarketClient).Init(opts.Host),
		account:   new(cli.AccountClient).Init(opts.AccessKey, opts.SecretKey, opts.Host),
		order:     new(cli.OrderClient).Init(opts.AccessKey, opts.SecretKey, opts.Host),
		wallet:    new(cli.WalletClient).Init(opts.AccessKey, opts.SecretKey, opts.Host),
		lastSleep: time.Now().UnixNano(),
	}
}

func (h *Huobi) GetType() string {
	return h.opts.Type
}

func (h *Huobi) GetName() string {
	return h.opts.Name
}

func (h *Huobi) AutoSleep() {
	now := time.Now().UnixNano()
	limit := h.opts.Limit

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

// GetAccountInfo
// spot：现货账户， margin：逐仓杠杆账户，
// otc：OTC 账户，point：点卡账户，super-margin：全仓杠杆账户,
// investment: C2C杠杆借出账户, borrow: C2C杠杆借入账户
func (h *Huobi) GetAccountInfo() []modacc.AccountInfo {
	rsp, err := h.account.GetAccountInfo()
	if err != nil {
		log.Err("get account info failed, %v", err)
	}
	return rsp
}

func (h *Huobi) GetAccountBalance(id string) *modacc.AccountBalance {
	rsp, err := h.account.GetAccountBalance(id)
	if err != nil {
		log.Err("get account balance faield, %v", err)
	}
	return rsp
}

func (h *Huobi) PlaceOrder(accountID, symbol, orderType, amount, price string) *modord.PlaceOrderResponse {
	req := &modpr.PlaceOrderRequest{
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
	}
	return rsp
}

var _ Exchange = &Huobi{}
