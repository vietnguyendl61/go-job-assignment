package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"pricing-service/model"
	"pricing-service/repo"
)

type PriceHandler struct {
	priceRepo repo.PriceRepo
}

func NewPriceHandler(priceRepo repo.PriceRepo) PriceHandler {
	return PriceHandler{priceRepo: priceRepo}
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

	price := &model.Price{}
	err = json.Unmarshal(body, &price)
	if err != nil {
		log.Println(err)
	}

	err = h.priceRepo.CreatePrice(r.Context(), price)
	if err != nil {
		log.Println("Error when create price: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(price)
}
