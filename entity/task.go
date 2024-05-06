package entity

import "time"

type Task struct {
	Id int `gorm:"primaryKey"`
	Name string
	Description string
	AssigneeId int
	ReviewerId int
	UserId int
	StatusId int
	CategoryId int
	ProjectId int
	Assignee     *User `gorm:"foreignKey:AssigneeId"`
    User         *User `gorm:"foreignKey:UserId"`
	Category *Category
	StartDate time.Time
	DueDate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}