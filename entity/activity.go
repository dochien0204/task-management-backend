package entity

import "time"

type Activity struct {
	Id          int `gorm:"primaryKey"`
	Description string
	UserId int
	TaskId int
	User *User

	CreatedAt time.Time
	UpdatedAt time.Time
}