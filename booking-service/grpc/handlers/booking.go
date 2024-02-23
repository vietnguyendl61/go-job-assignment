package handlers

import (
	bookingGrpc "booking-service/grpc/pb/booking"
	"booking-service/repo"
	"context"
	"encoding/json"
)

type GRPCHandlers struct {
	bookingGrpc.UnsafeBookingGrpcServer
	jobRepo repo.JobRepo
}

func NewGRPCHandlers(jobRepo repo.JobRepo) GRPCHandlers {
	return GRPCHandlers{
		jobRepo: jobRepo,
	}
}

func (h GRPCHandlers) GetListJobByBookDate(
	ctx context.Context,
	request *bookingGrpc.GetListJobByBookDateRequest,
) (
	*bookingGrpc.GetListJobByBookDateResponse,
	error,
) {
	var err error
	messageResponse := &bookingGrpc.GetListJobByBookDateResponse{}
	result, err := h.jobRepo.GetListJobByBookDate(ctx, request.BookDate)
	if err != nil {
		return nil, err
	}

	tmp, err := json.Marshal(&result)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(tmp, &messageResponse.JobList)
	if err != nil {
		return nil, err
	}
	return messageResponse, nil
}
