package model

import (
	"database/sql/driver"
	"encoding/json"
)

type UserSetting struct {
	UserID                  uint          `gorm:"primaryKey"`
	FreeTimeSlots           JSONTimeSlots `gorm:"type:json"`
	NotificationPreferences JSONMap       `gorm:"type:json"`

	// 关联关系
	User User `gorm:"foreignKey:UserID"`
}

type TimeSlot struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

// 自定义类型，用于处理JSON字段
type JSONTimeSlots []TimeSlot

func (j JSONTimeSlots) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONTimeSlots) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, &j)
}

type JSONMap map[string]interface{}

func (j JSONMap) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONMap) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, &j)
}
