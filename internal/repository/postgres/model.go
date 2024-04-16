package postgres

import "time"

type news struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Img         *string   `db:"img"`
	CreatedAt   time.Time `db:"create_at"`
}
