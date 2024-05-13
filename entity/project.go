package entity

import "time"

type Project struct {
	Id          int `gorm:"primaryKey"`
	Name        string
	Description string
	Image       string
	StatusId int
	Status *Status

	CreatedAt time.Time
	UpdatedAt time.Time
}
