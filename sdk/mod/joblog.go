package mod

import "time"

type LogType string

const (
	LogTypeStart LogType = "START"
	LogTypeStop  LogType = "STOP"
	LogTypeInfo  LogType = "INFO"
	LogTypeError LogType = "ERROR"
)

type JobLog struct {
	ID        int64      `xorm:"'id' pk autoincr" json:"id"`
	UserID    int64      `xorm:"'user_id' index" json:"userId"`
	JobID     int64      `xorm:"'job_id' index" json:"jobId"`
	Type      LogType    `xorm:"'type' varchar(25)" json:"logType"`
	Content   string     `xorm:"'content' varchar(250)" json:"content"`
	CreatedAt time.Time  `xorm:"created" json:"createdAt"`
	UpdatedAt time.Time  `xorm:"updated" json:"updatedAt"`
	DeletedAt *time.Time `xorm:"deleted index" json:"-"`
}
