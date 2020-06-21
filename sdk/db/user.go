package db

import (
	"aquarium/sdk/mod"
	"aquarium/utils/log"

	"github.com/go-xorm/xorm"
)

func GetUserByID(sess *xorm.Session, id int64) (*mod.User, error) {
	user := &mod.User{}
	found, err := sess.ID(id).Get(user)
	if err != nil {
		log.Err("get user by id failed, %v", err)
		return nil, err
	}
	if !found {
		return nil, nil
	}
	return user, nil
}

func GetUser(sess *xorm.Session, username string) (*mod.User, error) {
	user := &mod.User{}
	found, err := sess.Where("username = ?", username).Get(user)
	if err != nil {
		log.Err("get user by username failed, %v", err)
		return nil, err
	}
	if !found {
		return nil, nil
	}
	return user, nil
}

func ListUser(sess *xorm.Session, id, level, size, page int64, order string) (int64, []*mod.User, error) {
	var users []*mod.User
	bean := &mod.User{}
	total, err := sess.Where("level < ? OR id = ?", level, id).Count(bean)
	if err != nil {
		log.Err("count user failed, %v", err)
		return 0, nil, err
	}
	start, limit := int((page-1)*size), int(size)
	err = sess.Where("level < ? OR id = ?", level, id).OrderBy(order).Limit(limit, start).Find(&users)
	if err != nil {
		log.Err("list user failed, %v", err)
		return 0, nil, err
	}
	return total, users, nil
}

func CreateUser(sess *xorm.Session, user *mod.User) error {
	_, err := sess.InsertOne(user)
	if err != nil {
		log.Errw("create user failed", err)
		return err
	}
	return nil
}
