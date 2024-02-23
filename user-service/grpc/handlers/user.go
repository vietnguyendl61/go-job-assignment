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

func (h GRPCHandlers) GetAllUserId(ctx context.Context, request *userGrpc.GetAllHelperIdRequest) (*userGrpc.GetAllHelperIdResponse, error) {
	var (
		err          error
		listHelperId []string
		response     *userGrpc.GetAllHelperIdResponse
	)

	listHelperId, err = h.userRepo.GetListHelperId(ctx)
	if err != nil {
		return nil, err
	}

	response = &userGrpc.GetAllHelperIdResponse{
		ListHelperId: listHelperId,
	}
	return response, nil
}
