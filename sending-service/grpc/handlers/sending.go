package handlers

import (
	sendingGrpc "sending-service/grpc/pb/sending"
	"sending-service/repo"
)

type GRPCHandlers struct {
	sendingGrpc.UnimplementedPricingGrpcServer
	jobAssignmentRepo repo.JobAssignmentRepo
}

func NewGRPCHandlers(jobAssignmentRepo repo.JobAssignmentRepo) GRPCHandlers {
	return GRPCHandlers{
		jobAssignmentRepo: jobAssignmentRepo,
	}
}
