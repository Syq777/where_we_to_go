package model

import (
	"time"
)

type TaskCompletionRecord struct {
	RecordID        uint      `gorm:"primaryKey;autoIncrement"`
	TaskID          uint      `gorm:"not null"`
	UserID          uint      `gorm:"not null"`
	CompletedAt     time.Time `gorm:"autoCreateTime"`
	ActualTimeSpent *int      `gorm:"default:null"` // 单位：分钟

	// 关联关系
	User User `gorm:"foreignKey:UserID"`
	Task Task `gorm:"foreignKey:TaskID"`
}
