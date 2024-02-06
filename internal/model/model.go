package model

import "time"

type Model struct {
	ID        int `query:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
