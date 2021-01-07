package sdk

import (
	"context"

	"github.com/frankffenn/aquarium/config"
	"github.com/frankffenn/aquarium/sdk/db"
	"github.com/frankffenn/aquarium/sdk/mod"
)

func ListJob(ctx context.Context, ids []interface{}, size, page int64, order string) (int64, []*mod.Job, error) {
	return db.ListJob(config.Session(), ids, size, page, order)
}

func GetJobByID(ctx context.Context, id int64) (*mod.Job, error) {
	return db.GetJobByID(config.Session(), id)
}

func AddJob(ctx context.Context, Job *mod.Job) error {
	return db.AddJob(config.Session(), Job)
}

func UpdateJob(ctx context.Context, Job *mod.Job) error {
	return db.UpdateJob(config.Session(), Job)
}

func DeleteJob(ctx context.Context, ids []int64) error {
	return db.DeleteJob(config.Session(), ids)
}
