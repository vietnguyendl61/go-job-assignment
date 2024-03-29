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
	userGrpcHandlers "user-service/grpc/handlers"
	userGrpc "user-service/grpc/pb/user"
	"user-service/handlers"
	"user-service/model"
	"user-service/repo"
)

func main() {
	err := godotenv.Load()
	db := InitDB()
	if err != nil {
		log.Fatalln("Error when init db: " + err.Error())
	}
	MigrateDB(db)

	userRepo := repo.NewUserRepo(db)

	userHandlerGrpc := userGrpcHandlers.NewGRPCHandlers(userRepo)
	userHandler := handlers.NewUserHandler(userRepo)

	router := mux.NewRouter()
	router.HandleFunc("/user/register", userHandler.Register).Methods(http.MethodPost)
	router.HandleFunc("/user/login", userHandler.Login).Methods(http.MethodPost)

	go StartGRPCServer(userHandlerGrpc)

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
		&model.User{},
	}

	for _, m := range models {
		err := db.AutoMigrate(m)
		if err != nil {
			log.Println("Error when migrate: " + err.Error())
			return
		}
	}
}

func StartGRPCServer(handleGRPC userGrpcHandlers.GRPCHandlers) {
	var err error

	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("GRPC_PORT")))
	if err != nil {
		log.Fatalf("failed to listen GRPC: %v", err)
	}

	grpcServer := grpc.NewServer()
	userGrpc.RegisterUserGrpcServer(grpcServer, handleGRPC)

	log.Printf("Start listening GRPC server on port %s", os.Getenv("GRPC_PORT"))
	if err := grpcServer.Serve(grpcListener); err != nil {
		log.Fatalf("failed to listen GRPC: %v", err)
	}

	grpcServer.Stop()
}
