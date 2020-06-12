package mod

import "time"

// User
type User struct {
	ID        int64      `xorm:"'id' pk"`
	Username  string     `xorm: "'username' varchar(25) unique_index"`
	Password  string     `xorm:" not null"`
	Guid      string     `json:"'guid'"`
	Level     int64      `xorm:"'level'"`
	IsBanned  bool       `xorm:"'is_banned'`
	CreatedAt time.Time  `xorm:"created"`
	UpdatedAt time.Time  `xorm:"updated"`
	DeletedAt *time.Time `xorm:"-" deleted`
}
