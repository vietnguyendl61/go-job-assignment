package handlers

import (
	baseGrpc "booking-service/grpc"
	sendingGrpc "booking-service/grpc/pb/sending"
	"booking-service/model"
	"context"
	"os"
)

type SendingGrpcHandlers struct {
}

func NewSendingGrpcHandlers() SendingGrpcHandlers {
	return SendingGrpcHandlers{}
}

func (h SendingGrpcHandlers) clientPricingGrpc() (sendingGrpc.PricingGrpcClient, error) {
	conn, err := baseGrpc.ConnectGRPC(os.Getenv("PRICING_GRPC_HOST"), os.Getenv("PRICING_GRPC_PORT"))
	if err != nil {
		return nil, err
	}

	return sendingGrpc.NewPricingGrpcClient(conn), nil
}

func (h SendingGrpcHandlers) CreateJobAssignment(ctx context.Context, request model.CreateJobRequest) (*sendingGrpc.CreateJobAssignmentResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, baseGrpc.Timeout)
	defer cancel()
	grpcRequest := &sendingGrpc.CreateJobAssignmentRequest{
		JobId:     request.JobId.String(),
		ListJobId: request.ListJobId,
	}
	if request.CreatorId != nil {
		grpcRequest.CreatorId = request.CreatorId.String()
	}

	if request.CreatorId != nil {
		grpcRequest.CreatorId = request.CreatorId.String()
	}
	client, err := h.clientPricingGrpc()
	if err != nil {
		return nil, err
	}
	return client.CreateJobAssignment(ctx, grpcRequest)
}
