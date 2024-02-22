package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sending-service/model"
	"sending-service/repo"
)

type JobAssignmentHandler struct {
	jobAssignmentRepo repo.JobAssignmentRepo
}

func NewJobAssignmentHandler(jobAssignmentRepo repo.JobAssignmentRepo) JobAssignmentHandler {
	return JobAssignmentHandler{jobAssignmentRepo: jobAssignmentRepo}
}

func (h JobAssignmentHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			log.Println("Error when close request body: " + err.Error())
		}
	}()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	jobAssignment := &model.JobAssignment{}
	err = json.Unmarshal(body, &jobAssignment)
	if err != nil {
		log.Println(err)
	}

	result, err := h.jobAssignmentRepo.CreateJobAssignment(r.Context(), jobAssignment)
	if err != nil {
		log.Println("Error when create jobAssignment: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
