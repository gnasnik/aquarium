package trader

import (
	"github.com/frankffenn/aquarium/api"
	"github.com/frankffenn/aquarium/sdk/mod"
	"github.com/robertkrimen/otto"
)

type Global struct {
	*mod.Job
	Ctx    *otto.Otto
	ex     api.Exchange
	tasks  []task
	execed bool
	log    chan *mod.JobLog
}

type task struct {
	fn   otto.Value
	args []interface{}
}
