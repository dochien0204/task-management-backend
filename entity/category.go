package entity

import "time"

type Category struct {
	Id          int `gorm:"primaryKey"`
	Code        string
	Name        string
	Type        string
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
}