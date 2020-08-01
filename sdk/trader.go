package sdk

import (
	"context"

	"github.com/frankffenn/aquarium/config"
	"github.com/frankffenn/aquarium/sdk/db"
	"github.com/frankffenn/aquarium/sdk/mod"
)

func ListTrader(ctx context.Context, userID, algorithmID int64) ([]*mod.Trader, error) {
	sess := config.Session()
	traders, err := db.ListTrader(sess, userID, algorithmID)
	if err != nil {
		return nil, err
	}

	for i, x := range traders {
		exchange, err := db.GetExchangeFromTraderID(sess, x.ID)
		if err != nil {
			return nil, err
		}
		traders[i].Exchanges = append(traders[i].Exchanges, exchange)
	}

	return traders, nil
}

func GetTraderByID(ctx context.Context, id int64) (*mod.Trader, error) {
	return db.GetTraderByID(config.Session(), id)
}

func AddTrader(ctx context.Context, trader *mod.Trader) error {
	tx := config.Session()
	if err := tx.Begin(); err != nil {
		return err
	}

	if err := db.AddTrader(tx, trader); err != nil {
		tx.Rollback()
		return err
	}

	for _, x := range trader.Exchanges {
		traderExchange := &mod.TraderExchange{
			TraderID:   trader.ID,
			ExchangeID: x.ID,
		}
		if err := db.AddTraderExchange(tx, traderExchange); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func UpdateTrader(ctx context.Context, trader *mod.Trader) error {
	return db.UpdateTrader(config.Session(), trader)
}

func DeleteTrader(ctx context.Context, ids []int64) error {
	return db.DeleteTrader(config.Session(), ids)
}

func GetTraderExchangesByTraderID(id int64) ([]*mod.TraderExchange, error) {
	sess := config.Session()
	traderExchanges, err := db.GetTraderExchangeByTraderID(sess, id)
	if err != nil {
		return nil, err
	}
	for i, r := range traderExchanges {
		exchange, err := db.GetExchangeByID(sess, r.ExchangeID)
		if err != nil {
			return nil, err
		}

		traderExchanges[i].Exchange = *exchange
	}
	return traderExchanges, nil
}

func AddTraderExchange(ctx context.Context, t *mod.TraderExchange) error {
	return db.AddTraderExchange(config.Session(), t)
}
