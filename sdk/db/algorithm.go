package db

import (
	"aquarium/sdk/mod"
	"aquarium/utils/log"

	"github.com/go-xorm/xorm"
)

func ListAlgorithm(sess *xorm.Session, uid int64, size, page int64, order string) (int64, []*mod.Algorithm, error) {
	var out []*mod.Algorithm
	bean := &mod.Algorithm{}
	total, err := sess.Where("user_id = ?", uid).Count(bean)
	if err != nil {
		log.Err("count algorithm failed, %v", err)
		return 0, nil, err
	}
	start, limit := int((page-1)*size), int(size)
	err = sess.Where("user_id = ?", uid).OrderBy(order).Limit(limit, start).Find(&out)
	if err != nil {
		log.Err("list algorithm failed, %v", err)
		return 0, nil, err
	}
	return total, out, nil
}

func GetAlgorithmByID(sess *xorm.Session, id int64) (*mod.Algorithm, error) {
	out := &mod.Algorithm{ID: id}
	found, err := sess.Get(out)
	if err != nil {
		log.Err("get algorithm by id failed, %v", err)
		return nil, err
	}

	if !found {
		return nil, nil
	}

	return out, nil
}

func AddAlgorithm(sess *xorm.Session, Algorithm *mod.Algorithm) error {
	_, err := sess.InsertOne(Algorithm)
	if err != nil {
		log.Err("add algorithm failed, %v", err)
		return err
	}
	return nil
}

func UpdateAlgorithm(sess *xorm.Session, Algorithm *mod.Algorithm) error {
	cond := &mod.Algorithm{ID: Algorithm.ID}
	_, err := sess.Update(Algorithm, cond)
	if err != nil {
		log.Err("update algorithm failed, %v", err)
		return err
	}
	return nil
}

func DeleteAlgorithm(sess *xorm.Session, ids []int64) error {
	Algorithm := &mod.Algorithm{}
	_, err := sess.In("id", ids).Delete(Algorithm)
	if err != nil {
		log.Err("delete algorithm failed, %v", err)
		return err
	}
	return nil
}
