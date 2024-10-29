package model

import (
	"time"
)

type User struct {
	UserID       uint      `gorm:"primaryKey;autoIncrement"`
	Username     string    `gorm:"type:varchar(50);not null"`
	Email        string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	PasswordHash string    `gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`

	// 关联关系
	Tasks             []Task                 `gorm:"foreignKey:UserID"`
	UserSetting       *UserSetting           `gorm:"foreignKey:UserID"`
	CompletionRecords []TaskCompletionRecord `gorm:"foreignKey:UserID"`
	DailyTodoLists    []DailyTodoList        `gorm:"foreignKey:UserID"`
}
