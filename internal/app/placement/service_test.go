package placement

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/micro-it-freelance/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func NewTestPlacementService() Service {
	db, err := sqlx.Connect("pgx", fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=disable",
		config.DB.Name, config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port))
	if err != nil {
		panic(err)
	}

	return NewPlacementService(
		NewPlacementRepository(db),
	)
}
