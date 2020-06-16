package mod

import "time"

// User
type User struct {
	ID        int64      `xorm:"'id' pk autoincr"`
	Username  string     `xorm:"'username' varchar(25) unique "`
	Password  string     `xorm:"'password' not null"`
	Guid      string     `json:"'guid'"`
	Level     int64      `xorm:"'level'"`
	IsBanned  bool       `xorm:"'is_banned'`
	CreatedAt time.Time  `xorm:"created"`
	UpdatedAt time.Time  `xorm:"updated" `
	DeletedAt *time.Time `xorm:"-" deleted json:"omitempty"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) ToPlain() *PlainUser {
	return &PlainUser{
		ID:        u.ID,
		Username:  u.Username,
		Guid:      u.Guid,
		Level:     u.Level,
		IsBanned:  u.IsBanned,
		CreatedAt: u.CreatedAt,
	}
}

type PlainUser struct {
	ID        int64     `xorm:"'id' pk autoincr"`
	Username  string    `xorm:"'username' varchar(25) unique "`
	Guid      string    `json:"'guid'"`
	Level     int64     `xorm:"'level'"`
	IsBanned  bool      `xorm:"'is_banned'`
	CreatedAt time.Time `xorm:"created"`
}
