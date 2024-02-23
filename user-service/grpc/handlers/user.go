package handlers

import (
	userGrpc "user-service/grpc/pb/user"
	"user-service/repo"
)

type GRPCHandlers struct {
	userGrpc.UnimplementedPricingGrpcServer
	userRepo repo.UserRepo
}

func NewGRPCHandlers(userRepo repo.UserRepo) GRPCHandlers {
	return GRPCHandlers{
		userRepo: userRepo,
	}
}
