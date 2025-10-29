package entity

import "time"

type User struct {
	ID        string
	Email     string
	Password  string
	CreatedAt time.Time
}

type AuthToken struct {
	UserID string
	Token  string
	Expire time.Time
}
