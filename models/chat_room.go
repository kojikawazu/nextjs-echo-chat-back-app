package models

import (
	"time"
)

type ChatRooms struct {
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Users     []MiniAuthUsers `json:"users"`
}
