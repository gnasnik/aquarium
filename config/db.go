package config

import (
	_ "github.com/go-sql-driver/mysql"

	"aquarium/sdk/mod"
	adb "aquarium/utils/db"

	"github.com/go-xorm/xorm"
)

var (
	DB *xorm.Engine
)

func initDB() error {
	var err error
	DB, err = adb.OpenDB("mysql", Configs.DBURL)
	if err != nil {
		return err
	}

	DB.Sync2(new(mod.User))

	return nil
}

func Session() *xorm.Session {
	sess := DB.NewSession()
	return sess
}
