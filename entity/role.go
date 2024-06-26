package entity

import "time"

type Role struct {
	Id          int `gorm:"primaryKey"`
	Name        string
	Type        string
	Code        string
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
}
