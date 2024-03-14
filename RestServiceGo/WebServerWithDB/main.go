package main

import (
	"database-example/controller"
	"database-example/model"
	"database-example/repo"
	"database-example/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := "user=postgres password=super dbname=soa-tours host=localhost port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	database.AutoMigrate(&model.Tour{})
	return database
}


func startServer(controller *controller.TourController) {
	router := mux.NewRouter().StrictSlash(true)

	//TOURS
	//router.HandleFunc("/tours/{id}", controller.Get).Methods("GET")
	router.HandleFunc("/tours", controller.Create).Methods("POST")
	router.HandleFunc("/tours/authors/{id}", controller.GetAll).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8087", router))
}
func main() {

	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	tourRepository := &repo.TourRepository{DatabaseConnection: database}
	tourService := &service.TourService{TourRepo: tourRepository}
	tourController := &controller.TourController{TourService: tourService}

	startServer(tourController)
}
