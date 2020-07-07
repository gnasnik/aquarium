package sdk

import (
	"context"

	"github.com/frankffenn/aquarium/config"
	"github.com/frankffenn/aquarium/sdk/db"
	"github.com/frankffenn/aquarium/sdk/mod"
)

func GetUserByID(ctx context.Context, id int64) (*mod.User, error) {
	return db.GetUserByID(config.Session(), id)
}

func GetUser(ctx context.Context, username string) (*mod.User, error) {
	return db.GetUser(config.Session(), username)
}

func ListUser(ctx context.Context, id, level, size, page int64, order string) (int64, []*mod.User, error) {
	return db.ListUser(config.Session(), id, level, size, page, order)
}

func CreateUser(ctx context.Context, user *mod.User) error {
	return db.CreateUser(config.Session(), user)
}
