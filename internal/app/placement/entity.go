package placement

import (
	"time"

	core_entity "github.com/micro-it-freelance/core/entity"
	"github.com/micro-it-freelance/protoc/out/placement_service"
)

type Placement struct {
	ID         int64                           `db:"id" json:"db"`
	TelegramID int64                           `db:"telegram_id" json:"telegram_id"`
	Typ        placement_service.PlacementType `db:"typ" json:"typ"`
	Title      string                          `db:"title" json:"title"`
	Content    string                          `db:"content" json:"content"`
	CreatedAt  time.Time                       `db:"created_at" json:"created_at"`

	core_entity.ImplementedEntity
}
