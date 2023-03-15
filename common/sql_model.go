package common

import "time"

type SQLModel struct {
	ID int `json:"id"`
	// FakeID int
	Status    int        `json:"status"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
