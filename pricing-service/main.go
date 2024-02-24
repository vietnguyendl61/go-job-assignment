package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
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
	mongoClient, err := InitMongoDB()
	if err != nil {
		log.Fatalln("Error when init db: " + err.Error())
	}

	priceRepo := repo.NewPriceRepo(mongoClient)

	handlerGrpc := pricingGrpcHandlers.NewGRPCHandlers(priceRepo)
	priceHandler := handlers.NewPriceHandler(priceRepo)

	router := mux.NewRouter()
	router.HandleFunc("/price/get-list", priceHandler.GetPrice).Methods(http.MethodGet)

	go StartGRPCServer(handlerGrpc)

	log.Println("API is running in port: " + os.Getenv("PORT"))
	err = http.ListenAndServe(":"+os.Getenv("PORT"), router)
	if err != nil {
		log.Fatalln("Error: " + err.Error())
	}
}

func InitMongoDB() (*mongo.Client, error) {
	credential := options.Credential{
		Username: os.Getenv("MONGO_USER_NAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
	}

	clientOptions := options.Client().ApplyURI("mongodb://" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT")).SetAuth(credential)
	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
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
