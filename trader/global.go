package trader

import (
	"github.com/frankffenn/aquarium/api"
	"github.com/frankffenn/aquarium/sdk/mod"
	"github.com/robertkrimen/otto"
)

type Global struct {
	*mod.Trader
	Ctx       *otto.Otto
	es        []api.Exchange
	tasks     []task
	execed    bool
	statusLog string
}

type task struct {
	fn   otto.Value
	args []interface{}
}
