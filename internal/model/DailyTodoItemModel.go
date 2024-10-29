package model

import (
	"time"
)

type DailyTodoItem struct {
	ItemID        uint       `gorm:"primaryKey;autoIncrement"`
	ListID        uint       `gorm:"not null"`
	TaskID        uint       `gorm:"not null"`
	Status        string     `gorm:"type:varchar(20);default:'pending'"`
	ScheduledTime *time.Time `gorm:"default:null"`

	// 关联关系
	DailyTodoList DailyTodoList `gorm:"foreignKey:ListID"`
	Task          Task          `gorm:"foreignKey:TaskID"`
}
