package handlers

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sending-service/repo"
	"sending-service/utils"
)

type JobAssignmentHandler struct {
	jobAssignmentRepo repo.JobAssignmentRepo
}

func NewJobAssignmentHandler(jobAssignmentRepo repo.JobAssignmentRepo) JobAssignmentHandler {
	return JobAssignmentHandler{jobAssignmentRepo: jobAssignmentRepo}
}

func (h JobAssignmentHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	Id := mux.Vars(r)["job_id"]

	UUID, err := uuid.Parse(Id)
	if err != nil {
		log.Println("Error when parse id: " + err.Error())
		utils.ErrorResponse(w, http.StatusBadRequest, "Error when parse id: "+err.Error())
		return
	}

	job, err := h.jobAssignmentRepo.GetOneByJobId(r.Context(), UUID.String())
	if err != nil {
		log.Println("Error when get job: " + err.Error())
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error when get job: "+err.Error())
		return
	}

	utils.ResponseWithData(w, http.StatusOK, job)
	return
}
