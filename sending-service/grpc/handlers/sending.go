package handlers

import (
	sendingGrpc "sending-service/grpc/pb/sending"
	"sending-service/repo"
)

type GRPCHandlers struct {
	sendingGrpc.UnimplementedPricingGrpcServer
	jobAssignmentRepo repo.JobAssignmentRepo
	userHandlerGrpc   UserGrpcHandlers
}

func NewGRPCHandlers(
	jobAssignmentRepo repo.JobAssignmentRepo,
	userHandlerGrpc UserGrpcHandlers,
) GRPCHandlers {
	return GRPCHandlers{
		jobAssignmentRepo: jobAssignmentRepo,
		userHandlerGrpc:   userHandlerGrpc,
	}
}
