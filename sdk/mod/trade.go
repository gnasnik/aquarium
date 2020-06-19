package mod

import "time"

// Trader struct
type Trader struct {
	ID          int64      `xorm:"'id' pk autoincr" json:"id"`
	UserID      int64      `xorm:"'user_id' index" json:"userId"`
	AlgorithmID int64      `xorm:"'algorithm_id' index" json:"algorithmId"`
	Name        string     `xorm:"'name' varchar(200)" json:"name"`
	Environment string     `xorm:"'environment' text" json:"environment"`
	LastRunAt   time.Time  `xorm:"'last_run_at'" json:"lastRunAt"`
	CreatedAt   time.Time  `xorm:"created" json:"createdAt"`
	UpdatedAt   time.Time  `xorm:"updated" json:"updatedAt"`
	DeletedAt   *time.Time `xorm:"deleted" sql:"index" json:"-"`

	Exchanges []Exchange `xorm:"-" json:"exchanges"`
	Status    int64      `xorm:"-" json:"status"`
	Algorithm Algorithm  `xorm:"-" json:"algorithm"`
}
