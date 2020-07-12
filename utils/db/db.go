package db

import (
	"time"

	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

func OpenDB(driver, url string) (*xorm.Engine, error) {
	db, err := xorm.NewEngine(driver, url)
	if err != nil {
		return nil, err
	}

	// if err := KeepAlive(db); err != nil {
	// 	return nil, err
	// }

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return nil, err
	}

	db.DatabaseTZ = location

	return db, err
}

func KeepAlive(db *xorm.Engine) error {
	if err := db.Ping(); err != nil {
		return err
	}

	db.Logger().SetLevel(core.LOG_ERR)

	go func() {
		for {
			time.Sleep(time.Minute)
			db.Ping()
		}
	}()

	return nil
}
