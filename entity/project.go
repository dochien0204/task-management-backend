package entity

import "time"

type Project struct {
	Id          int `gorm:"primaryKey"`
	Name        string
	Description string
	Image       string
	StatusId int
	Status *Status
	UserId int 
	User *User

	CreatedAt time.Time
	UpdatedAt time.Time
}
