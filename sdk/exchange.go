package sdk

import (
	"aquarium/config"
	"aquarium/sdk/db"
	"aquarium/sdk/mod"
	"context"
)

func ListExchange(ctx context.Context, ids []interface{}, size, page int64, order string) (int64, []*mod.Exchange, error) {
	return db.ListExchange(config.Session(), ids, size, page, order)
}

func GetExchangeByID(ctx context.Context, id int64) (*mod.Exchange, error) {
	return db.GetExchangeByID(config.Session(), id)
}

func AddExchange(ctx context.Context, exchange *mod.Exchange) error {
	return db.AddExchange(config.Session(), exchange)
}

func UpdateExchange(ctx context.Context, exchange *mod.Exchange) error {
	return db.UpdateExchange(config.Session(), exchange)
}
