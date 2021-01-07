package mod

import "time"

type Job struct {
	ID          int64      `xorm:"'id' pk autoincr" json:"id"`
	UserID      int64      `xorm:"'user_id' index" json:"userId"`
	Name        string     `xorm:"'name' varchar(50)" json:"name"`
	AlgorithmID int64      `xorm:"'algorithm_id'" json:"algorithm_id"`
	ExchangeID  int64      `xorm:"'exchange_id'" json:"exchange_id"`
	Running     bool       `xorm:"'running'" json:"running"`
	CreatedAt   time.Time  `xorm:"created" json:"createdAt"`
	UpdatedAt   time.Time  `xorm:"updated" json:"updatedAt"`
	DeletedAt   *time.Time `xorm:"deleted index" json:"-"`
}
