package entity

import "time"

type TaskDocument struct {
	Id int `gorm:"primaryKey"`
	File string
	Name string
	TaskId int

	CreatedAt time.Time
	UpdatedAt time.Time
}