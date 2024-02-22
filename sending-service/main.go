package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"sending-service/handlers"
	"sending-service/repo"
)

func main() {
	err := godotenv.Load()
	db := InitDB()
	if err != nil {
		log.Fatalln("Error when init db: " + err.Error())
	}

	migrationHandler := handlers.NewMigrationHandler(db)

	jobAssignment := repo.NewJobAssignmentRepo(db)
	jobHandler := handlers.NewJobAssignmentHandler(jobAssignment)

	router := mux.NewRouter()
	router.HandleFunc("/migration", migrationHandler.Migrate).Methods(http.MethodGet)

	router.HandleFunc("/job-assignment/create", jobHandler.Create).Methods(http.MethodPost)

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
