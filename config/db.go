package config

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/frankffenn/aquarium/sdk/mod"
	adb "github.com/frankffenn/aquarium/utils/db"

	"github.com/go-xorm/xorm"
)

var (
	_defaultEngine *xorm.Engine
)

func initDB() error {
	var err error

	_defaultEngine, err = adb.OpenDB("mysql", Configs.DBURL)
	if err != nil {
		return err
	}

	_defaultEngine.Sync2(
		new(mod.User),
		new(mod.Exchange),
		new(mod.Algorithm),
		new(mod.Trader),
		new(mod.TraderExchange),
	)
	return nil
}

func Session() *xorm.Session {
	sess := _defaultEngine.NewSession()
	return sess
}
