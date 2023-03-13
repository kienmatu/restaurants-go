package common

import "time"

type SQLModel struct {
	ID int
	// FakeID int
	Status    int        `json:"status" gorm:"column:status,default:1"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at,autoCreateTime"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at,autoUpdateTime"`
}
