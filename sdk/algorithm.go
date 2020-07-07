package sdk

import (
	"github.com/frankffenn/aquarium/config"
	"github.com/frankffenn/aquarium/sdk/db"
	"github.com/frankffenn/aquarium/sdk/mod"
	"context"
)

func ListAlgorithm(ctx context.Context, uid, size, page int64, order string) (int64, []*mod.Algorithm, error) {
	return db.ListAlgorithm(config.Session(), uid, size, page, order)
}

func GetAlgorithmByID(ctx context.Context, id int64) (*mod.Algorithm, error) {
	return db.GetAlgorithmByID(config.Session(), id)
}

func AddAlgorithm(ctx context.Context, Algorithm *mod.Algorithm) error {
	return db.AddAlgorithm(config.Session(), Algorithm)
}

func UpdateAlgorithm(ctx context.Context, Algorithm *mod.Algorithm) error {
	return db.UpdateAlgorithm(config.Session(), Algorithm)
}

func DeleteAlgorithm(ctx context.Context, ids []int64) error {
	return db.DeleteAlgorithm(config.Session(), ids)
}
