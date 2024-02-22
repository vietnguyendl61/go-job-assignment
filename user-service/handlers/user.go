package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"user-service/model"
	"user-service/repo"
	"user-service/utils"
)

type UserHandler struct {
	userRepo repo.UserRepo
}

func NewUserHandler(userRepo repo.UserRepo) UserHandler {
	return UserHandler{userRepo: userRepo}
}

func (h UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			log.Println("Error when close request body: " + err.Error())
		}
	}()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error when parse request: " + err.Error())
		utils.ErrorResponse(w, http.StatusBadRequest, "Error when parse request: "+err.Error())
	}

	user := &model.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error when parse request: " + err.Error())
		utils.ErrorResponse(w, http.StatusBadRequest, "Error when parse request: "+err.Error())
	}

	result, err := h.userRepo.CreateUser(r.Context(), user)
	if err != nil {
		log.Println("Error when create user: " + err.Error())
		utils.ErrorResponse(w, http.StatusBadRequest, "Error when create user: "+err.Error())
	}
	utils.ResponseWithData(w, http.StatusCreated, result)
}

func (h UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			log.Println("Error when close request body: " + err.Error())
		}
	}()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error when parse request: " + err.Error())
		utils.ErrorResponse(w, http.StatusBadRequest, "Error when parse request: "+err.Error())
	}

	request := model.LoginRequest{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Println("Error when parse request: " + err.Error())
		utils.ErrorResponse(w, http.StatusBadRequest, "Error when parse request: "+err.Error())
	}

	result, err := h.userRepo.GetUserByUserNameAndPassword(r.Context(), request)
	if err != nil {
		log.Println("Error when login: " + err.Error())
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error when login: "+err.Error())
	}
	if result == nil {
		log.Println("Wrong user name or password")
		utils.ErrorResponse(w, http.StatusBadRequest, "Wrong user name or password")
	}
	utils.ResponseWithData(w, http.StatusOK, result.ID)
}
