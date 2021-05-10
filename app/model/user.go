package model

import (
	"time"
)

type User struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Mobile       int64     `json:"mobile" xorm:"not null default 0 comment('手机号') unique BIGINT(11)"`
	Email        string    `json:"email" xorm:"not null default '' comment('邮箱') unique VARCHAR(64)"`
	Nickname     string    `json:"nickname" xorm:"not null default '' VARCHAR(32)"`
	Sex          int       `json:"sex" xorm:"not null default 0 comment('性别') TINYINT(1)"`
	Avatar       int       `json:"avatar" xorm:"not null default 0 comment('头像') INT(11)"`
	AuthKey      string    `json:"-",xorm:"not null default '' VARCHAR(32)"`
	PasswordHash string    `json:"-",xorm:"not null default '' VARCHAR(255)"`
	Status       int       `json:"status" xorm:"not null default 0 TINYINT(1)"`
	UserType     int       `json:"user_type" xorm:"not null default 0 comment('0.普通用户 1.管理员') TINYINT(1)"`
	CreatedAt    time.Time `json:"created_at" xorm:"created TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at" xorm:"updated TIMESTAMP"`
}

func (t *User) TableName() string {
	return "user"
}
