package placement

import "time"

type PlacementEnum string

const (
	PlacementOffer PlacementEnum = "offer"
	PlacementOrder PlacementEnum = "order"
)

type Placement struct {
	ID int64 `db:"id" json:"db"`

	TelegramID int64     `db:"telegram_id" json:"telegram_id"`
	Username   string    `db:"username" json:"username"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`

	Type PlacementEnum `db:"type" json:"type"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Contacts    string `db:"contacts" json:"contacts"`
}
