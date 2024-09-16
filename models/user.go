package models

type User struct {
	UserID   int
	Username string
	Password string
}

type UserDetail struct {
	UserID   int    `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
}
