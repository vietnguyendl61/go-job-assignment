package handlers

import (
	"context"
	"github.com/google/uuid"
	"os"
	baseGrpc "pricing-service/grpc"
	bookingGrpc "pricing-service/grpc/pb/booking"
)

type BookingGrpcHandlers struct {
}

func NewBookingGrpcHandlers() BookingGrpcHandlers {
	return BookingGrpcHandlers{}
}

func (h BookingGrpcHandlers) clientPricingGrpc() (bookingGrpc.BookingGrpcClient, error) {
	conn, err := baseGrpc.ConnectGRPC(os.Getenv("BOOKING_GRPC_HOST"), os.Getenv("BOOKING_GRPC_PORT"))
	if err != nil {
		return nil, err
	}

	return bookingGrpc.NewBookingGrpcClient(conn), nil
}

func (h BookingGrpcHandlers) GetListJobByBookDate(
	ctx context.Context,
	creatorId uuid.UUID,
	bookDate string,
) (
	*bookingGrpc.GetListJobByBookDateResponse,
	error,
) {
	ctx, cancel := context.WithTimeout(ctx, baseGrpc.Timeout)
	defer cancel()
	grpcRequest := &bookingGrpc.GetListJobByBookDateRequest{
		CreatorId: creatorId.String(),
		BookDate:  bookDate,
	}

	client, err := h.clientPricingGrpc()
	if err != nil {
		return nil, err
	}
	return client.GetListJobByBookDate(ctx, grpcRequest)
}
