package entities

import "time"

type PriorityLevel string

const (
	Low      PriorityLevel = "low"
	Mid      PriorityLevel = "mid"
	High     PriorityLevel = "high"
	VeryHigh PriorityLevel = "very-high"
)

type Todo struct {
	ID              int           `json:"id" gorm:"column:id; primaryKey"`
	ActivityGroupId int           `json:"activity_group_id" gorm:"column:activity_group_id"`
	Title           string        `json:"title" gorm:"column:title"`
	IsActive        bool          `json:"is_active" gorm:"column:is_active"`
	Priority        PriorityLevel `json:"priority" gorm:"type:ENUM('low', 'mid', 'high', 'very-high');default:'very-high'"`
	CreatedAt       time.Time     `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt       time.Time     `json:"updatedAt" gorm:"column:updatedAt"`
}
