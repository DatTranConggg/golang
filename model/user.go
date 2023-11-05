package model

import "time"

type User struct {
	UserId    string    `json:"user_id" db:"user_id, omitempty"`
	FullName  string    `json:"full_name" db:"full_name, omitempty"`
	Email     string    `json:"email" db:"email, omitempty"`
	Password  string    `json:"-" db:"password, omitempty"`
	Role      string    `json:"-" db:"role, omitempty"`
	CreatedAt time.Time `json:"-" db:"created_at, omitempty"`
	UpdatedAt time.Time `json:"-" db:"updated_at, omitempty"`
	Token     string    `json:"token" db:"token, omitempty"`
}
