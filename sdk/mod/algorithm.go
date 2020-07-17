package mod

import "time"

type Algorithm struct {
	ID          int64      `xorm:"'id' pk autoincr" json:"id"`
	UserID      int64      `xorm:"'user_id' index" json:"userId"`
	Name        string     `xorm:"'name' varchar(200)" json:"name"`
	Description string     `xorm:"'description' text" json:"description"`
	Script      string     `xorm:"'script' text" json:"script"`
	EvnDefault  string     `xorm:"'evn_default' text" json:"evnDefault"`
	CreatedAt   time.Time  `xorm:"created" json:"createdAt"`
	UpdatedAt   time.Time  `xorm:"updated" json:"updatedAt"`
	DeletedAt   *time.Time `xorm:"deleted" json:"-"`

	Traders []*Trader `xorm:"-" json:"traders"`
}
