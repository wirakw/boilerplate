package entities

import "time"

type Activity struct {
	ID        int       `json:"id" gorm:"column:activity_id; primaryKey"`
	Title     string    `json:"title" gorm:"column:title"`
	Email     string    `json:"email" gorm:"column:email"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}
