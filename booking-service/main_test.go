package main

import (
	bookingGrpc "booking-service/grpc/pb/booking"
	"booking-service/handlers"
	grpcHandler "booking-service/mock/grpc/handlers"
	"github.com/google/uuid"

	"booking-service/repo"
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	"net/http/httptest"
	"testing"
)

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)`

func TestInitDB(t *testing.T) {
	_ = godotenv.Load()
	db := InitDB()

	t.Run("connect to database", func(t *testing.T) {
		if err := db.Exec(tableCreationQuery).Error; err != nil {
			t.Errorf("Error when connect to db: " + err.Error())
		}
		db.Exec("DELETE FROM products")
		db.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
	})
}

type mock struct {
	GetListJobByBookDate func() (*bookingGrpc.GetListJobByBookDateResponse, error)
}

func TestCreateJobAndVerifyInDatabase(t *testing.T) {
	// Create a new router instance
	r := mux.NewRouter()

	db := InitDB()

	jobRepo := repo.NewJobRepo(db)

	messageResponse := &bookingGrpc.GetListJobByBookDateResponse{}
	messageResponse.JobList = []*bookingGrpc.Job{
		{
			Id:          uuid.New().String(),
			BookDate:    "2024-04-23T11:17:05.360024+00:00",
			Description: "Something",
		},
	}

	mockPricing := &mock{GetListJobByBookDate: func() (*bookingGrpc.GetListJobByBookDateResponse, error) { return messageResponse, nil }}
	sendHandlerGrpc := grpcHandler.NewSendingGrpcHandlers()

	jobHandler := handlers.NewJobHandler(jobRepo, mockPricing, sendHandlerGrpc)

	// Register your handler function with the router
	r.HandleFunc("/job/create", jobHandler.Create).Methods(http.MethodPost)

	// Prepare sample data
	requestBody := map[string]interface{}{
		"book_date":   "2024-04-23T11:17:05.360024+00:00",
		"description": "Sample job description",
		"price":       99.99,
	}

	// Marshal the request body to JSON
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, "/job/create", bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Serve the request using the router
	r.ServeHTTP(rr, req)

	// Check the status code of the response
	if rr.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %v want %v", rr.Code, http.StatusOK)
	}

	t.Errorf("Something")
	//
	//// Now let's fetch the inserted record from the database
	//// Use your GetOneJob function from your JobRepo
	//insertedJob, err := jobRepo.GetOneJob(context.Background(), requestBody.JobId.String())
	//if err != nil {
	//	t.Fatalf("error fetching inserted job from database: %v", err)
	//}
	//
	//// Compare the values of the inserted record with the values from the API request
	//if insertedJob.CreatorId != requestBody.CreatorId.String() {
	//	t.Errorf("expected CreatorId %s but got %s", requestBody.CreatorId.String(), insertedJob.CreatorId)
	//}
	//if insertedJob.Description != requestBody.Description {
	//	t.Errorf("expected Description %s but got %s", requestBody.Description, insertedJob.Description)
	//}
	// Add similar checks for other fields as needed
}
