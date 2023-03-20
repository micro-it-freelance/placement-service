package placement

import (
	"github.com/jmoiron/sqlx"
)

type Repository interface {

}

type AccountRepository struct {
	db *sqlx.DB
}

func NewPlacementRepository(db *sqlx.DB) Repository {
	return &AccountRepository{
		db: db,
	}
}

