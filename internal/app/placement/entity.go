package placement

import (
	"time"

	"github.com/micro-it-freelance/protoc/out/placement_service"
)

type Placement struct {
	ID         int64                           `db:"id" json:"db"`
	TelegramID int64                           `db:"telegram_id" json:"telegram_id"`
	Username   string                          `db:"username" json:"username"`
	CreatedAt  time.Time                       `db:"created_at" json:"created_at"`
	Typ        placement_service.PlacementType `db:"typ" json:"typ"`
	Title      string                          `db:"title" json:"title"`
	Content    string                          `db:"content" json:"content"`
}
