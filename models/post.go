package models

import "time"

// 记得内存对齐
type Post struct {
	PostID      int       `json:"post_id" db:"post_id"`
	AuthorID    int       `json:"author_id" db:"author_id"`
	Status      int       `json:"status" db:"status"`
	CommunityID int       `json:"community_id" db:"community_id" binding:"required"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}
