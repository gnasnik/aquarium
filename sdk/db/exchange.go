package db

import (
	"aquarium/sdk/mod"
	"aquarium/utils/log"

	"github.com/go-xorm/xorm"
)

func ListExchange(sess *xorm.Session, ids []interface{}, size, page int64, order string) (int64, []*mod.Exchange, error) {
	var out []*mod.Exchange
	bean := &mod.Exchange{}
	total, err := sess.In("user_id", ids...).Count(bean)
	if err != nil {
		log.Err("count exchange failed, %v", err)
		return 0, nil, err
	}
	start, limit := int((page-1)*size), int(size)
	err = sess.In("user_id", ids...).OrderBy(order).Limit(limit, start).Find(&out)
	if err != nil {
		log.Err("list exchange failed, %v", err)
		return 0, nil, err
	}
	return total, out, nil
}

func GetExchangeByID(sess *xorm.Session, id int64) (*mod.Exchange, error) {
	out := &mod.Exchange{ID: id}
	found, err := sess.Get(out)
	if err != nil {
		log.Err("get exchange by id failed, %v", err)
		return nil, err
	}

	if !found {
		return nil, nil
	}

	return out, nil
}

func AddExchange(sess *xorm.Session, exchange *mod.Exchange) error {
	_, err := sess.InsertOne(exchange)
	if err != nil {
		log.Err("add exchange failed, %v", err)
		return err
	}
	return nil
}

func UpdateExchange(sess *xorm.Session, exchange *mod.Exchange) error {
	cond := &mod.Exchange{ID: exchange.ID}
	_, err := sess.Update(exchange, cond)
	if err != nil {
		log.Err("update exchange failed, %v", err)
		return err
	}
	return nil
}

func DeleteExchange(sess *xorm.Session, ids []int64) error {
	exchange := &mod.Exchange{}
	_, err := sess.In("id", ids).Delete(exchange)
	if err != nil {
		log.Err("delete exchange failed, %v", err)
		return err
	}
	return nil
}

func AddTraderExchange(sess *xorm.Session, t *mod.TraderExchange) error {
	_, err := sess.InsertOne(t)
	if err != nil {
		log.Err("add trader exchange failed, %v", err)
		return err
	}
	return nil
}

func GetExchangeFromTraderID(sess *xorm.Session, traderID int64) (*mod.Exchange, error) {
	sql := `SELECT e.* FROM exchanges e, trader_exchanges r WHERE r.trader_id
	= ? AND e.id = r.exchange_id`

	var out *mod.Exchange
	found, err := sess.Sql(sql, traderID).Get(out)
	if err != nil {
		log.Err("get exchange by trader id failed, %v", err)
		return nil, err
	}

	if !found {
		return nil, nil
	}

	return out, nil
}
