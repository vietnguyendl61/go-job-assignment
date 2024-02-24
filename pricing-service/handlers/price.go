package handlers

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"net/http"
	grpcHandler "pricing-service/grpc/handlers"
	"pricing-service/model"
	"pricing-service/repo"
	"pricing-service/utils"
)

type PriceHandler struct {
	priceRepo          repo.PriceRepo
	bookingHandlerGrpc grpcHandler.BookingGrpcHandlers
}

func NewPriceHandler(priceRepo repo.PriceRepo) PriceHandler {
	return PriceHandler{priceRepo: priceRepo}
}

func (h PriceHandler) GetPrice(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("x-user-id")
	if userId == "" {
		log.Println("Missing x-user-id")
		utils.ErrorResponse(w, http.StatusBadRequest, "You must login first")
		return
	}

	userUUID, err := uuid.Parse(userId)
	if err != nil {
		log.Println("Error when parse user id: " + err.Error())
		utils.ErrorResponse(w, http.StatusBadRequest, "Error when parse user id: "+err.Error())
		return
	}

	date := r.URL.Query().Get("date")

	res, err := h.bookingHandlerGrpc.GetListJobByBookDate(r.Context(), userUUID, date)
	if err != nil {
		log.Println("Error when get list job: " + err.Error())
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error when get list job: "+err.Error())
		return
	}

	var result []*model.Price
	for _, v := range res.JobList {
		price, err := h.priceRepo.GetPriceByJobId(r.Context(), v.Id)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Error when get price: " + err.Error())
			utils.ErrorResponse(w, http.StatusInternalServerError, "Error when parse user id: "+err.Error())
			return
		}
		result = append(result, price)
	}

	// Send a 201 created response
	utils.ResponseWithData(w, http.StatusOK, result)
}
