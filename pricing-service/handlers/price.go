package handlers

import (
	"booking-service/model"
	"booking-service/repo"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type PriceHandler struct {
	jobRepo repo.PriceRepo
}

func NewPriceHandler(jobRepo repo.PriceRepo) PriceHandler {
	return PriceHandler{jobRepo: jobRepo}
}

func (h PriceHandler) Create(w http.ResponseWriter, r *http.Request) {
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

	job := &model.Price{}
	err = json.Unmarshal(body, &job)
	if err != nil {
		log.Println(err)
	}

	result, err := h.jobRepo.CreatePrice(r.Context(), job)
	if err != nil {
		log.Println("Error when create job: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
