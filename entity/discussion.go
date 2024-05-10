package entity

import "time"

type Discussion struct {
	Id          int `gorm:"primaryKey"`
	Comment string
	UserId int
	TaskId int
	User *User

	CreatedAt time.Time
	UpdatedAt time.Time
}