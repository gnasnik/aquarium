package mod

import "time"

type Exchange struct {
	ID        int64      `xorm:"'id' pk autoincr" json:"id"`
	UserID    int64      `xorm:"'user_id' index" json:"userId"`
	Name      string     `xorm:"'name' varchar(50)" json:"name"`
	Type      string     `xorm:"'type' varchar(50)" json:"type"`
	AccessKey string     `xorm:"'access_key' varchar(200)" json:"accessKey"`
	SecretKey string     `xorm:"'secret_key' varchar(200)" json:"secretKey"`
	CreatedAt time.Time  `xorm:"created" json:"createdAt"`
	UpdatedAt time.Time  `xorm:"updated" json:"updatedAt"`
	DeletedAt *time.Time `xorm:"deleted index" json:"-"`
}
