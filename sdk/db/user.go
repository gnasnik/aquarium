package db

import (
	"aquarium/sdk/mod"

	"github.com/go-xorm/xorm"
)

func GetUserByID(sess *xorm.Session, id int64) (user *mod.User, err error) {
	_, err = sess.ID(id).Get(user)
	return
}

func GetUser(sess *xorm.Session, username string) (user *mod.User, err error) {
	_, err = sess.Where("username = ?", username).Get(user)
	return
}

func ListUser(sess *xorm.Session, id, level, size, page int64, order string) (total int64, users []*mod.User, err error) {
	total, err = sess.Where("level < ? OR id = ?", level, id).Count(&users)
	if err != nil {
		return
	}
	start, limit := int((page-1)*size), int(size)
	err = sess.Where("level < ? OR id = ?", level, id).OrderBy(order).Limit(limit, start).Find(&users)
	return
}

func CreateUser(sess *xorm.Session, user *mod.User) error {
	_, err := sess.InsertOne(sess, user)
	return err
}
