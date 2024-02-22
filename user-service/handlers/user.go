package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"user-service/model"
	"user-service/repo"
)

type UserHandler struct {
	userRepo repo.UserRepo
}

func NewUserHandler(userRepo repo.UserRepo) UserHandler {
	return UserHandler{userRepo: userRepo}
}

func (h UserHandler) Create(w http.ResponseWriter, r *http.Request) {
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

	user := &model.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println(err)
	}

	result, err := h.userRepo.CreateUser(r.Context(), user)
	if err != nil {
		log.Println("Error when create user: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
