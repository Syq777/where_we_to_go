package model

import (
	"time"
)

type DailyTodoList struct {
	ListID      uint      `gorm:"primaryKey;autoIncrement"`
	UserID      uint      `gorm:"not null"`
	Date        time.Time `gorm:"type:date;not null"`
	GeneratedAt time.Time `gorm:"autoCreateTime"`

	// 关联关系
	User      User            `gorm:"foreignKey:UserID"`
	TodoItems []DailyTodoItem `gorm:"foreignKey:ListID"`
}
