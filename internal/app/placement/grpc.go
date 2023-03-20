package placement

import (
	"github.com/micro-it-freelance/protoc/out/account_service"
)

type PlacementGRPCHandler struct {
	service Service
	account_service.UnimplementedAccountServiceServer
}

func NewAccountGRPCHandler(s Service) account_service.AccountServiceServer {
	return &PlacementGRPCHandler{
		service: s,
	}
}
