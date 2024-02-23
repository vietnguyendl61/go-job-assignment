package handlers

import (
	baseGrpc "booking-service/grpc"
	pricingGrpc "booking-service/grpc/pb/pricing"
	"booking-service/model"
	"context"
	"os"
)

type PriceGrpcHandlers struct {
}

func NewPriceGrpcHandlers() PriceGrpcHandlers {
	return PriceGrpcHandlers{}
}

func (h PriceGrpcHandlers) clientPricingGrpc() (pricingGrpc.PricingGrpcClient, error) {
	conn, err := baseGrpc.ConnectGRPC(os.Getenv("PRICING_GRPC_HOST"), os.Getenv("PRICING_GRPC_PORT"))
	if err != nil {
		return nil, err
	}

	return pricingGrpc.NewPricingGrpcClient(conn), nil
}

func (h PriceGrpcHandlers) CreatePrice(ctx context.Context, request model.CreateJobRequest) (*pricingGrpc.CreatePriceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, baseGrpc.Timeout)
	defer cancel()
	grpcRequest := &pricingGrpc.CreatePriceRequest{
		JobId: request.JobId.String(),
		Price: request.Price,
	}

	if request.CreatorId != nil {
		grpcRequest.CreatorId = request.CreatorId.String()
	}
	client, err := h.clientPricingGrpc()
	if err != nil {
		return nil, err
	}
	return client.CreatePrice(ctx, grpcRequest)
}
