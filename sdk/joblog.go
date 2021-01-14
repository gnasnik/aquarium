package sdk

import (
	"context"

	"github.com/frankffenn/aquarium/config"
	"github.com/frankffenn/aquarium/sdk/db"
	"github.com/frankffenn/aquarium/sdk/mod"
)

func ListJobLog(ctx context.Context, id interface{}, logType string, size, page int64, order string) (int64, []*mod.JobLog, error) {
	return db.ListJobLog(config.Session(), id, logType, size, page, order)
}

func AddJobLog(ctx context.Context, Job *mod.JobLog) error {
	return db.AddJobLog(config.Session(), Job)
}
