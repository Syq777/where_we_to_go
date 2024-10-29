package model

import (
	"time"
)

type Task struct {
	TaskID          uint      `gorm:"primaryKey;autoIncrement"`
	UserID          uint      `gorm:"not null"`
	ParentTaskID    *uint     `gorm:"default:null"`
	TaskName        string    `gorm:"type:varchar(255);not null"`
	TaskDescription string    `gorm:"type:text"`
	EstimatedTime   *int      `gorm:"default:null"` // 单位：分钟
	Priority        int       `gorm:"default:0"`
	Status          string    `gorm:"type:varchar(20);default:'pending'"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`

	// 关联关系
	User              User                   `gorm:"foreignKey:UserID"`
	ParentTask        *Task                  `gorm:"foreignKey:ParentTaskID"`
	SubTasks          []Task                 `gorm:"foreignKey:ParentTaskID"`
	TodoItems         []DailyTodoItem        `gorm:"foreignKey:TaskID"`
	CompletionRecords []TaskCompletionRecord `gorm:"foreignKey:TaskID"`
}
