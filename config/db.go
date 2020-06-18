package config

import (
	_ "github.com/go-sql-driver/mysql"

	"aquarium/sdk/mod"
	adb "aquarium/utils/db"

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
	)
	return nil
}

func Session() *xorm.Session {
	sess := _defaultEngine.NewSession()
	return sess
}
