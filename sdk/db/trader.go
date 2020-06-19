package db

import (
	"aquarium/sdk/mod"
	"aquarium/utils/log"

	"github.com/go-xorm/xorm"
)

func ListTrader(sess *xorm.Session, userID, algorithmID int64) ([]*mod.Trader, error) {
	var out []*mod.Trader
	err := sess.Where("user_id = ?", userID).And("algorithm_id = ?", algorithmID).Find(&out)
	if err != nil {
		log.Err("list trader failed, %v", err)
		return nil, err
	}
	return out, nil
}

func GetTraderByID(sess *xorm.Session, id int64) (*mod.Trader, error) {
	out := &mod.Trader{ID: id}
	found, err := sess.Get(out)
	if err != nil {
		log.Err("get trader by id failed, %v", err)
		return nil, err
	}

	if !found {
		return nil, nil
	}

	return out, nil
}

func AddTrader(sess *xorm.Session, Trader *mod.Trader) error {
	_, err := sess.InsertOne(Trader)
	if err != nil {
		log.Err("add trader failed, %v", err)
		return err
	}
	return nil
}

func UpdateTrader(sess *xorm.Session, Trader *mod.Trader) error {
	cond := &mod.Trader{ID: Trader.ID}
	_, err := sess.Update(Trader, cond)
	if err != nil {
		log.Err("update trader failed, %v", err)
		return err
	}
	return nil
}

func DeleteTrader(sess *xorm.Session, ids []int64) error {
	Trader := &mod.Trader{}
	_, err := sess.In("id", ids).Delete(Trader)
	if err != nil {
		log.Err("delete trader failed, %v", err)
		return err
	}
	return nil
}

func GetTraderExchangeByTraderID(sess *xorm.Session, traderID int64) ([]*mod.TraderExchange, error) {
	var out []*mod.TraderExchange
	err := sess.Where("trader_id = ?", traderID).Find(&out)
	if err != nil {
		log.Err("get traderExchange by trader_id failed, %v", err)
		return nil, err
	}

	return out, nil
}
