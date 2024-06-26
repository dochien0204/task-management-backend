package entity

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	Id int `gorm:"primaryKey"`
	Name string
	Description string
	AssigneeId *int
	ReviewerId *int
	UserId int
	StatusId int
	Status *Status
	CategoryId *int
	ProjectId int
	Assignee     *User `gorm:"foreignKey:AssigneeId"`
    User         *User `gorm:"foreignKey:UserId"`
	Reviewer *User `gorm:"foreignKey:ReviewerId"`
	Category *Category
	Documents []*TaskDocument `gorm:"foreignKey:TaskId"`
	StartDate *time.Time
	DueDate *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`
}