package model

import "time"

type User struct {
	ID           int64     `json:"id"`
	Username     string    `json:"username" validate:"required,max=50"`
	Email        string    `json:"email" validate:"required,email,max=100"`
	PasswordHash string    `json:"password_hash" validate:"required,max=255"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserCredential struct {
	Username string `json:"username" validate:"required,max=30"` // big int di postgresql
	Password string `json:"password" validate:"required"`
}
