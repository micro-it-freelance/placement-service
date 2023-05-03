package placement

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	core_db "github.com/micro-it-freelance/core/database"
)

// var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func NewTestPlacementService() Service {
	// connect to database
	db := core_db.NewDBConnection()

	return NewPlacementService(
		NewPlacementRepository(db),
	)
}
