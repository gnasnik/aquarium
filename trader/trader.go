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
	if GetTraderStatus(id) {
		return stop(id)
	}
	return run(id)
}

func GetTraderStatus(id int64) bool {
	if t, ok := Executor[id]; ok && t != nil {
		return t.Running
	}
	return false
}

func run(id int64) error {
	job, err := initialize(id)
	if err != nil {
		log.Err("initialize trader failed, %v", err)
		return err
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errw("recover", "err", err)
			}
		}()
		job.LastRunAt = time.Now()
		job.Running = true
		if job.Algorithm.Script == "" {
			log.Errw("empty script", "aligorithm id", job.Algorithm.ID)
			return
		}
		if _, err := job.Ctx.Run(job.Algorithm.Script); err != nil {
			log.Err("run script failed, %v", err)
			return
		}
		main, err := job.Ctx.Get("main")
		if err != nil {
			log.Err("Can not get the main function")
			return
		}
		if _, err := main.Call(main); err != nil {
			log.Err("call main function failed,%v", err)
			return
		}
	}()
	Executor[job.ID] = job
	return nil
}

func stop(id int64) error {
	if t, ok := Executor[id]; !ok || t == nil {
		return xerrors.New("Can not found the Trader")
	}
	Executor[id].Ctx.Interrupt <- func() { panic(errHalt) }
	Executor[id].Job.Running = false
	return nil
}

func initialize(id int64) (*Global, error) {
	if t := Executor[id]; t != nil && t.Running {
		return nil, nil
	}

	ctx := context.Background()
	job, err := sdk.GetJobByID(ctx, id)
	if err != nil {
		log.Err("get trader by id failed,%v", err)
		return nil, err
	}

	if job.AlgorithmID <= 0 {
		return nil, xerrors.New("Please select a algorithm")
	}

	job.Algorithm, err = sdk.GetAlgorithmByID(ctx, job.AlgorithmID)
	if err != nil {
		log.Err("get algorithm by id failed,%v", err)
		return nil, err
	}

	e, err := sdk.GetExchangeByID(ctx, job.ExchangeID)
	if err != nil {
		log.Err("get traderexchange by id failed,%v", err)
		return nil, err
	}

	ex := createExchange(
		comm.ExchangeType(e.Type),
		api.JobID(job.ID),
		api.Name(e.Name),
		api.Type(e.Type),
		api.AccessKey(e.AccessKey),
		api.SecretKey(e.SecretKey),
	)

	global := &Global{
		Job:   job,
		tasks: make([]task, 0),
		Ctx:   otto.New(),
		ex:    ex,
	}
	for _, c := range comm.Consts {
		global.Ctx.Set(c, c)
	}
	global.Ctx.Interrupt = make(chan func(), 1)
	global.Ctx.Set("Global", global)
	global.Ctx.Set("G", global)
	global.Ctx.Set("Exchange", global.ex)
	global.Ctx.Set("E", global.ex)
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
