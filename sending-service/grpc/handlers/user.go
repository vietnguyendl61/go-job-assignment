package handlers

import (
	"context"
	"os"
	baseGrpc "sending-service/grpc"
	userGrpc "sending-service/grpc/pb/user"
)

type UserGrpcHandlers struct {
}

func NewUserGrpcHandlers() UserGrpcHandlers {
	return UserGrpcHandlers{}
}

func (h UserGrpcHandlers) clientUserGrpc() (userGrpc.UserGrpcClient, error) {
	conn, err := baseGrpc.ConnectGRPC(os.Getenv("USER_GRPC_HOST"), os.Getenv("USER_GRPC_PORT"))
	if err != nil {
		return nil, err
	}
	return userGrpc.NewUserGrpcClient(conn), nil
}

func (h UserGrpcHandlers) GetAllHelperId(ctx context.Context) (*userGrpc.GetAllHelperIdResponse, error) {
	var (
		err         error
		grpcRequest *userGrpc.GetAllHelperIdRequest
	)
	ctx, cancel := context.WithTimeout(ctx, baseGrpc.Timeout)
	defer cancel()

	client, err := h.clientUserGrpc()
	if err != nil {
		return nil, err
	}

	return client.GetAllHelperId(ctx, grpcRequest)
}
