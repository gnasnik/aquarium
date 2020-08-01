package trader

import (
	"context"
	"time"

	"github.com/frankffenn/aquarium/api"
	"github.com/frankffenn/aquarium/comm"
	"github.com/frankffenn/aquarium/sdk"
	"github.com/frankffenn/aquarium/utils/log"
	"github.com/robertkrimen/otto"
	"golang.org/x/xerrors"
)

var (
	errHalt  = xerrors.New("HALT")
	Executor = make(map[int64]*Global)
)

func Switch(id int64) error {
	if GetTraderStatus(id) > 0 {
		return stop(id)
	}
	return run(id)
}

func GetTraderStatus(id int64) int64 {
	if t, ok := Executor[id]; ok && t != nil {
		return t.Status
	}
	return 0
}

func run(id int64) error {
	trader, err := initialize(id)
	if err != nil {
		return err
	}

	go func() {
		trader.LastRunAt = time.Now()
		trader.Status = 1
		if _, err := trader.Ctx.Run(trader.Algorithm.Script); err != nil {
			log.Err("run script failed, %v", err)
			return
		}
		main, err := trader.Ctx.Get("main")
		if err != nil {
			log.Err("Can not get the main function")
			return
		}
		if _, err := main.Call(main); err != nil {
			log.Err("call main function failed,%v", err)
			return
		}
	}()
	Executor[trader.ID] = trader
	return nil
}

func stop(id int64) error {
	if t, ok := Executor[id]; !ok || t == nil {
		return xerrors.New("Can not found the Trader")
	}
	Executor[id].Ctx.Interrupt <- func() { panic(errHalt) }
	return nil
}

func initialize(id int64) (*Global, error) {
	if t := Executor[id]; t != nil && t.Status > 0 {
		return nil, nil
	}

	ctx := context.Background()
	trader, err := sdk.GetTraderByID(ctx, id)
	if err != nil {
		log.Err("get trader by id failed,%v", err)
		return nil, err
	}

	if trader.AlgorithmID <= 0 {
		return nil, xerrors.New("Please select a algorithm")
	}

	trader.Algorithm, err = sdk.GetAlgorithmByID(ctx, trader.AlgorithmID)
	if err != nil {
		log.Err("get algorithm by id failed,%v", err)
		return nil, err
	}

	es, err := sdk.GetTraderExchangesByTraderID(trader.ID)
	if err != nil {
		log.Err("get traderexchange by id failed,%v", err)
		return nil, err
	}

	global := &Global{
		Trader: trader,
		tasks:  make([]task, 0),
		Ctx:    otto.New(),
	}
	for _, c := range comm.Consts {
		global.Ctx.Set(c, c)
	}
	for _, e := range es {
		ex := createExchange(
			comm.ExchangeType(e.Type),
			api.TraderID(trader.ID),
			api.Name(e.Name),
			api.Type(e.Type),
			api.AccessKey(e.AccessKey),
			api.SecretKey(e.SecretKey),
		)
		global.es = append(global.es, ex)
	}
	if len(global.es) == 0 {
		return nil, xerrors.New("Please add at least one exchange")
	}
	global.Ctx.Interrupt = make(chan func(), 1)
	global.Ctx.Set("Global", global)
	global.Ctx.Set("G", global)
	global.Ctx.Set("Exchange", global.es[0])
	global.Ctx.Set("E", global.es[0])
	global.Ctx.Set("Exchanges", global.es)
	global.Ctx.Set("Es", global.es)
	return global, nil
}

func createExchange(t comm.ExchangeType, opts ...api.Option) api.Exchange {
	switch t {
	case comm.Huobi:
		return api.NewHuobi(opts...)
	default:
	}
	return nil
}

func clean(userID int64) {
	for _, t := range Executor {
		if t != nil && t.UserID == userID {
			stop(t.ID)
		}
	}
}
