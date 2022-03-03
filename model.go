package utils

import "time"

type Model struct {
	ID        string    `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
