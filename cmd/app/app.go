package main

import (
	_ "github.com/jackc/pgx/v5/stdlib"

	core_db "github.com/micro-it-freelance/core/database"
	core_grpc "github.com/micro-it-freelance/core/grpc"

	"github.com/micro-it-freelance/placement-service/internal/app/placement"
	"github.com/micro-it-freelance/protoc/out/placement_service"
	"google.golang.org/grpc"
)

func main() {
	// connect to database
	db := core_db.NewDB()

	// add listener
	listener := core_grpc.NewGRPCListener()

	// create grpc server
	GRPCServer := grpc.NewServer()
	placement_service.RegisterPlacementServiceServer(GRPCServer,
		placement.NewPlacementGRPCHandler(
			placement.NewPlacementService(
				placement.NewPlacementRepository(db),
			),
		))

	// serve
	core_grpc.Serve(GRPCServer, listener)
}
