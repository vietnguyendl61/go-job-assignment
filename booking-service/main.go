package main

import (
	grpcHandler "booking-service/grpc/handlers"
	bookingGrpc "booking-service/grpc/pb/booking"
	"booking-service/handlers"
	"booking-service/model"
	"booking-service/repo"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	db := InitDB()
	if err != nil {
		log.Fatalln("Error when init db: " + err.Error())
	}
	MigrateDB(db)

	jobRepo := repo.NewJobRepo(db)

	bookingHandlerGrpc := grpcHandler.NewGRPCHandlers(jobRepo)
	priceHandlerGrpc := grpcHandler.NewPriceGrpcHandlers()
	sendHandlerGrpc := grpcHandler.NewSendingGrpcHandlers()

	jobHandler := handlers.NewJobHandler(jobRepo, priceHandlerGrpc, sendHandlerGrpc)

	router := mux.NewRouter()

	router.HandleFunc("/job/create", jobHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/job/get-one/{id}", jobHandler.GetOne).Methods(http.MethodGet)

	go StartGRPCServer(bookingHandlerGrpc)
	log.Println("API is running in port: " + os.Getenv("PORT"))
	err = http.ListenAndServe(":"+os.Getenv("PORT"), router)
	if err != nil {
		log.Fatalln("Error: " + err.Error())
	}
}

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func MigrateDB(db *gorm.DB) {
	_ = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	models := []interface{}{
		&model.Job{},
	}

	for _, m := range models {
		err := db.AutoMigrate(m)
		if err != nil {
			log.Println("Error when migrate: " + err.Error())
			return
		}
	}
}

func StartGRPCServer(handleGRPC grpcHandler.GRPCHandlers) {
	var err error

	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("GRPC_PORT")))
	if err != nil {
		log.Fatalf("failed to listen GRPC: %v", err)
	}

	grpcServer := grpc.NewServer()
	bookingGrpc.RegisterBookingGrpcServer(grpcServer, handleGRPC)

	log.Printf("Start listening GRPC server on port %s", os.Getenv("GRPC_PORT"))
	if err := grpcServer.Serve(grpcListener); err != nil {
		log.Fatalf("failed to listen GRPC: %v", err)
	}

	grpcServer.Stop()
}
