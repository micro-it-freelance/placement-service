package placement

import (
	"github.com/micro-it-freelance/protoc/out/placement_service"
)

type PlacementGRPCHandler struct {
	service Service
	placement_service.UnimplementedPlacementServiceServer
}

func NewPlacementGRPCHandler(s Service) placement_service.PlacementServiceServer {
	return &PlacementGRPCHandler{
		service: s,
	}
}
