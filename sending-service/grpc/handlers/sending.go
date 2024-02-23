package handlers

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	sendingGrpc "sending-service/grpc/pb/sending"
	userGrpc "sending-service/grpc/pb/user"
	"sending-service/model"
	"sending-service/repo"
)

type GRPCHandlers struct {
	sendingGrpc.UnsafeSendingGrpcServer
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

func (h GRPCHandlers) CreateJobAssignment(ctx context.Context, request *sendingGrpc.CreateJobAssignmentRequest) (*sendingGrpc.CreateJobAssignmentResponse, error) {
	var (
		err           error
		userResponse  *userGrpc.GetAllHelperIdResponse
		response      *sendingGrpc.CreateJobAssignmentResponse
		creatorId     uuid.UUID
		jobId         uuid.UUID
		jobAssignment *model.JobAssignment
	)

	jobId, err = uuid.Parse(request.JobId)
	if err != nil {
		return nil, err
	}

	userResponse, err = h.userHandlerGrpc.GetAllHelperId(ctx)
	if err != nil {
		return nil, err
	}

	listHelperId, err := h.jobAssignmentRepo.GetListHelperIdByJobId(ctx, request.ListJobId)
	if err != nil {
		return nil, err
	}

	if len(listHelperId) == 0 && len(userResponse.ListHelperId) > 0 {
		userId, err := uuid.Parse(userResponse.ListHelperId[rand.Intn(len(userResponse.ListHelperId))])
		if err != nil {
			return nil, err
		}
		jobAssignment = &model.JobAssignment{
			HelperId: userId,
			JobId:    jobId,
		}
	} else if len(listHelperId) > 0 {
		res := elementsNotInSlice(userResponse.ListHelperId, listHelperId)
		if len(res) > 0 {
			userId, err := uuid.Parse(res[rand.Intn(len(res))])
			if err != nil {
				return nil, err
			}

			jobAssignment = &model.JobAssignment{
				HelperId: userId,
				JobId:    jobId,
			}
		} else {
			return nil, fmt.Errorf("No helper is available")
		}
	} else {
		return nil, fmt.Errorf("No helper is available")
	}

	if request.CreatorId != "" {
		creatorId, err = uuid.Parse(request.CreatorId)
		if err != nil {
			return nil, err
		}
		jobAssignment.BaseModel.CreatorID = &creatorId
	}
	jobAssignment.JobStatus = "Processing"

	err = h.jobAssignmentRepo.CreateJobAssignment(ctx, jobAssignment)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func elementsNotInSlice(slice1, slice2 []string) []string {
	slice2Map := make(map[string]bool)
	for _, value := range slice2 {
		slice2Map[value] = true
	}

	var result []string

	for _, value := range slice1 {
		if !slice2Map[value] {
			result = append(result, value)
		}
	}

	return result
}
