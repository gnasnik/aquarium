package db

import (
	"github.com/frankffenn/aquarium/sdk/mod"
	"github.com/frankffenn/aquarium/utils/log"

	"github.com/go-xorm/xorm"
)

func ListJob(sess *xorm.Session, ids []interface{}, size, page int64, order string) (int64, []*mod.Job, error) {
	var out []*mod.Job
	bean := &mod.Job{}
	total, err := sess.In("user_id", ids...).Count(bean)
	if err != nil {
		log.Err("count Job failed, %v", err)
		return 0, nil, err
	}
	start, limit := int((page-1)*size), int(size)
	err = sess.In("user_id", ids...).OrderBy(order).Limit(limit, start).Find(&out)
	if err != nil {
		log.Err("list Job failed, %v", err)
		return 0, nil, err
	}
	return total, out, nil
}

func GetJobByID(sess *xorm.Session, id int64) (*mod.Job, error) {
	out := &mod.Job{ID: id}
	found, err := sess.Get(out)
	if err != nil {
		log.Err("get Job by id failed, %v", err)
		return nil, err
	}

	if !found {
		return nil, nil
	}

	return out, nil
}

func AddJob(sess *xorm.Session, Job *mod.Job) error {
	_, err := sess.InsertOne(Job)
	if err != nil {
		log.Err("add Job failed, %v", err)
		return err
	}
	return nil
}

func UpdateJob(sess *xorm.Session, Job *mod.Job) error {
	cond := &mod.Job{ID: Job.ID}
	_, err := sess.Update(Job, cond)
	if err != nil {
		log.Err("update Job failed, %v", err)
		return err
	}
	return nil
}

func DeleteJob(sess *xorm.Session, ids []int64) error {
	Job := &mod.Job{}
	_, err := sess.In("id", ids).Delete(Job)
	if err != nil {
		log.Err("delete Job failed, %v", err)
		return err
	}
	return nil
}
