package handlers

import (
	"context"
	userGrpc "user-service/grpc/pb/user"
	"user-service/repo"
)

type GRPCHandlers struct {
	userGrpc.UnimplementedUserGrpcServer
	userRepo repo.UserRepo
}

func NewGRPCHandlers(userRepo repo.UserRepo) GRPCHandlers {
	return GRPCHandlers{
		userRepo: userRepo,
	}
}

func (h GRPCHandlers) GetAllUserId(ctx context.Context, request *userGrpc.GetAllUserIdRequest) (*userGrpc.GetAllUserIdResponse, error) {
	var (
		err        error
		listUserId []string
		response   *userGrpc.GetAllUserIdResponse
	)

	listUserId, err = h.userRepo.GetListUerId(ctx)
	if err != nil {
		return nil, err
	}

	response = &userGrpc.GetAllUserIdResponse{
		ListUserId: listUserId,
	}
	return response, nil
}
