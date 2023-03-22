package common

import "time"

type SQLModel struct {
	ID        int        `json:"-" gorm:"column:id;"`
	FakeID    *UID       `json:"id" gorm:"-"`
	Status    int        `json:"status" gorm:"column:status;default:1;"`
	CreatedAt *time.Time `json:"createdAt,omitempty" gorm:"column:createdAt;"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"column:updatedAt;"`
}

func (s *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(s.ID), dbType, 1)
	s.FakeID = &uid
}
