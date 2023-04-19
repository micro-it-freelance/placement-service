package placement

import (
	"github.com/jmoiron/sqlx"
)

type Repository interface {
}

type PlacementRepository struct {
	db *sqlx.DB
}

func NewPlacementRepository(db *sqlx.DB) Repository {
	return &PlacementRepository{
		db: db,
	}
}
