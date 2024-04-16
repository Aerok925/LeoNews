package models

import "time"

type News struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Img         *string   `json:"img"`
	CreatedAt   time.Time `json:"create_at"`
}
