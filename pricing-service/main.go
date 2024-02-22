package main

import (
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
	pricingGrpcHandlers "pricing-service/grpc/handlers"
	pricingGrpc "pricing-service/grpc/pb/pricing"
	"pricing-service/handlers"
	"pricing-service/repo"
)

func main() {
	err := godotenv.Load()
	db := InitDB()
	if err != nil {
		log.Fatalln("Error when init db: " + err.Error())
	}

	priceRepo := repo.NewPriceRepo(db)

	migrationHandler := handlers.NewMigrationHandler(db)
	handlerGrpc := pricingGrpcHandlers.NewGRPCHandlers(priceRepo)
	priceHandler := handlers.NewPriceHandler(priceRepo)

	router := mux.NewRouter()
	router.HandleFunc("/migration", migrationHandler.Migrate).Methods(http.MethodGet)

	router.HandleFunc("/price/create", priceHandler.Create).Methods(http.MethodPost)

	go StartGRPCServer(handlerGrpc)

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

func StartGRPCServer(handleGRPC pricingGrpcHandlers.GRPCHandlers) {
	var err error

	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("GRPC_PORT")))
	if err != nil {
		log.Fatalf("failed to listen GRPC: %v", err)
	}

	grpcServer := grpc.NewServer()
	pricingGrpc.RegisterPricingGrpcServer(grpcServer, handleGRPC)

	log.Printf("Start listening GRPC server on port %s", os.Getenv("GRPC_PORT"))
	if err := grpcServer.Serve(grpcListener); err != nil {
		log.Fatalf("failed to listen GRPC: %v", err)
	}

	grpcServer.Stop()
}
