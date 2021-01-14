package db

import (
	"github.com/frankffenn/aquarium/sdk/mod"
	"github.com/frankffenn/aquarium/utils/log"

	"github.com/go-xorm/xorm"
)

func ListJobLog(sess *xorm.Session, id interface{}, logType string, size, page int64, order string) (int64, []*mod.JobLog, error) {
	var out []*mod.JobLog
	bean := &mod.JobLog{}

	s1 := sess.Where("job_id = ?", id)
	if logType != "" {
		s1 = s1.And("type = ?", logType)
	}
	total, err := s1.Count(bean)
	if err != nil {
		log.Err("count Job log failed, %v", err)
		return 0, nil, err
	}

	sess = sess.Where("job_id = ?", id)
	if logType != "" {
		sess = sess.And("type = ?", logType)
	}
	start, limit := int((page-1)*size), int(size)
	err = sess.OrderBy(order).Limit(limit, start).Find(&out)
	if err != nil {
		log.Err("list Job log failed, %v", err)
		return 0, nil, err
	}
	return total, out, nil
}

func AddJobLog(sess *xorm.Session, job *mod.JobLog) error {
	_, err := sess.InsertOne(job)
	if err != nil {
		log.Err("add Job log failed, %v", err)
		return err
	}
	return nil
}
