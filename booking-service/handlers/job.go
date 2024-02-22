package handlers

import (
	"booking-service/model"
	"booking-service/repo"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type JobHandler struct {
	jobRepo *repo.JobRepo
}

func NewJobHandler(jobRepo *repo.JobRepo) *JobHandler {
	return &JobHandler{jobRepo: jobRepo}
}

func (h *JobHandler) Create(w http.ResponseWriter, r *http.Request) {
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

	job := &model.Job{}
	err = json.Unmarshal(body, &job)
	if err != nil {
		log.Println(err)
	}

	result, err := h.jobRepo.CreateJob(r.Context(), job)
	if err != nil {
		log.Println("Error when create job: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
