package handlers

import (
	grpcHandler "booking-service/grpc/handlers"
	"booking-service/model"
	"booking-service/repo"
	"booking-service/utils"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
)

type JobHandler struct {
	jobRepo          repo.JobRepo
	priceHandlerGrpc grpcHandler.PriceGrpcHandlers
	sendHandlerGrpc  grpcHandler.SendingGrpcHandlers
}

func NewJobHandler(
	jobRepo repo.JobRepo,
	priceHandlerGrpc grpcHandler.PriceGrpcHandlers,
	sendHandlerGrpc grpcHandler.SendingGrpcHandlers,
) JobHandler {
	return JobHandler{
		jobRepo:          jobRepo,
		priceHandlerGrpc: priceHandlerGrpc,
		sendHandlerGrpc:  sendHandlerGrpc,
	}
}

func (h JobHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error when parse request: " + err.Error())
		utils.ErrorResponse(w, http.StatusBadRequest, "Error when parse request: "+err.Error())
	}

	jobRequest := model.CreateJobRequest{}
	err = json.Unmarshal(body, &jobRequest)
	if err != nil {
		log.Println("Error when parse request: " + err.Error())
		utils.ErrorResponse(w, http.StatusBadRequest, "Error when parse request: "+err.Error())
		return
	}
	if jobRequest.Price <= 0 {
		log.Println("Invalid Price")
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid Price")
		return
	}
	if jobRequest.BookDate == nil {
		log.Println("Missing booking date")
		utils.ErrorResponse(w, http.StatusBadRequest, "Missing booking date")
		return
	}
	job := &model.Job{
		BookDate:    *jobRequest.BookDate,
		Description: jobRequest.Description,
	}
	userId := r.Header.Get("x-user-id")
	if userId != "" {
		userUUID, err := uuid.Parse(userId)
		if err != nil {
			log.Println("Error when parse user id: " + err.Error())
			utils.ErrorResponse(w, http.StatusBadRequest, "Error when parse user id: "+err.Error())
			return
		}
		jobRequest.CreatorId = &userUUID
		job.CreatorID = &userUUID
	}

	tx, cancel := h.jobRepo.CreateTx(r.Context())
	defer func() {
		tx.Rollback()
		err := r.Body.Close()
		if err != nil {
			log.Println("Error when close request body: " + err.Error())
		}
		cancel()
	}()
	tx = tx.Begin()
	err = h.jobRepo.CreateJob(tx, job)
	if err != nil {
		log.Println("Error when create job: " + err.Error())
		utils.ErrorResponse(w, http.StatusBadRequest, "Error when create job: "+err.Error())
		return
	}
	jobRequest.JobId = job.ID

	var listId []string
	listId, err = h.jobRepo.GetListJobIdByBookDate(r.Context(), job.BookDate)
	if err != nil {
		log.Println("Error when get list job id by book date: " + err.Error())
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error when get list job id by book date: "+err.Error())
		return
	}
	jobRequest.ListJobId = listId

	_, err = h.sendHandlerGrpc.CreateJobAssignment(tx.Statement.Context, jobRequest)
	if err != nil {
		log.Println("Error when create job assignment: " + err.Error())
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error when create job assignment: "+err.Error())
		return
	}

	_, err = h.priceHandlerGrpc.CreatePrice(tx.Statement.Context, jobRequest)
	if err != nil {
		log.Println("Error when create job: " + err.Error())
		utils.ErrorResponse(w, http.StatusBadRequest, "Error when create job: "+err.Error())
		return
	}

	tx.Commit()
	utils.ResponseWithData(w, http.StatusCreated, job)
	return
}
