package handlers

import (
	"context"
	pricingGrpc "pricing-service/grpc/pb/pricing"
)

type GRPCHandlers struct {
	pricingGrpc.UnimplementedPricingGrpcServer
}

func NewGRPCHandlers() GRPCHandlers {
	return GRPCHandlers{}
}

func (h GRPCHandlers) CreatePrice(ctx context.Context, request *pricingGrpc.CreatePriceRequest) (*pricingGrpc.CreatePriceResponse, error) {

	return nil, nil
}
