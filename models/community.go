package models

type Community struct {
	ID   int    `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}
