package models

import "time"

type AuthUsers struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MiniAuthUsers struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
