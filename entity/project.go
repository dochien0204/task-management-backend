package entity

import (
	"time"

	"gorm.io/gorm"
)

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
	DeleteAt gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

type ProjectTaskCount struct {
	Project Project `gorm:"embedded"`
	TaskCount int `gorm:"column:task_count"`
}

type ProjectMemberCount struct {
	Project Project `gorm:"embedded"`
	MemberCount int `gorm:"column:member_count"`
}