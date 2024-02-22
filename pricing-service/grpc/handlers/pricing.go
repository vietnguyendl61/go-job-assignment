package handlers

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	pricingGrpc "pricing-service/grpc/pb/pricing"
	"pricing-service/model"
	"pricing-service/repo"
)

type GRPCHandlers struct {
	pricingGrpc.UnimplementedPricingGrpcServer
	priceRepo repo.PriceRepo
}

func NewGRPCHandlers(priceRepo repo.PriceRepo) GRPCHandlers {
	return GRPCHandlers{
		priceRepo: priceRepo,
	}
}

func (h GRPCHandlers) CreatePrice(ctx context.Context, request *pricingGrpc.CreatePriceRequest) (*pricingGrpc.CreatePriceResponse, error) {
	var (
		err             error
		jobId           uuid.UUID
		creatorId       uuid.UUID
		price           *model.Price
		messageResponse *pricingGrpc.CreatePriceResponse
	)
	jobId, err = uuid.Parse(request.JobId)
	if err != nil {
		messageResponse = &pricingGrpc.CreatePriceResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Error when parse job id: " + err.Error(),
		}
		return messageResponse, err
	}

	if request.CreatorId != "" {
		creatorId, err = uuid.Parse(request.CreatorId)
		if err != nil {
			messageResponse = &pricingGrpc.CreatePriceResponse{
				StatusCode: http.StatusBadRequest,
				Message:    "Error when parse creator id: " + err.Error(),
			}
			return messageResponse, err
		}
		price.BaseModel.CreatorID = &creatorId
	}

	price = &model.Price{
		JobId: jobId,
		Price: request.Price,
	}

	err = h.priceRepo.CreatePrice(ctx, price)
	if err != nil {
		messageResponse = &pricingGrpc.CreatePriceResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error when create price: " + err.Error(),
		}
		return messageResponse, err
	}

	messageResponse = &pricingGrpc.CreatePriceResponse{
		StatusCode: http.StatusCreated,
		Message:    "Create price success",
	}
	return messageResponse, nil
}
