package mod

import "time"

type Job struct {
	ID          int64      `xorm:"'id' pk autoincr" json:"id"`
	UserID      int64      `xorm:"'user_id' index" json:"userId"`
	Name        string     `xorm:"'name' varchar(50)" json:"name"`
	Description string     `xorm:"'description' varchar(50)" json:"description"`
	AlgorithmID int64      `xorm:"'algorithm_id'" json:"algorithmId"`
	ExchangeID  int64      `xorm:"'exchange_id'" json:"exchangeId"`
	Running     bool       `xorm:"'running'" json:"running"`
	LastRunAt   time.Time  `xorm:"'last_run_at'" json:"lastRunAt"`
	CreatedAt   time.Time  `xorm:"created" json:"createdAt"`
	UpdatedAt   time.Time  `xorm:"updated" json:"updatedAt"`
	DeletedAt   *time.Time `xorm:"deleted index" json:"-"`

	Algorithm *Algorithm `xorm:"-" json:"algorithm"`
	Exchanges *Exchange  `xorm:"-" json:"exchange"`
}
