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
	database.AutoMigrate(&model.Equipment{})
	database.AutoMigrate(&model.KeyPoint{})
	database.AutoMigrate(&model.TourPreference{})

	return database
}

func startServer(tourController *controller.TourController,
	equipmentController *controller.EquipmentController,
	keyPointController *controller.KeyPointController,
	tourPreferenceController *controller.TourPreferenceController) {
	router := mux.NewRouter().StrictSlash(true)

	//TOURS
	//router.HandleFunc("/tours/{id}", controller.Get).Methods("GET")
	router.HandleFunc("/tours", tourController.Create).Methods("POST")
	router.HandleFunc("/tours/authors/{id}", tourController.GetAll).Methods("GET")
	router.HandleFunc("/tours/{id}", tourController.Get).Methods("GET")

	//equipment
	router.HandleFunc("/equipment", equipmentController.Create).Methods("POST")
	router.HandleFunc("/equipment", equipmentController.GetAll).Methods("GET")
	router.HandleFunc("/equipment/{id}", equipmentController.Get).Methods("GET")

	//key points
	router.HandleFunc("/key-points", keyPointController.Create).Methods("POST")
	router.HandleFunc("/tours/{tourId}/key-points", keyPointController.GetAll).Methods("GET")

	//tour preferences
	router.HandleFunc("/tour-preferences", tourPreferenceController.Create).Methods("POST")
	router.HandleFunc("/tourists/{userId}/tour-preference", tourPreferenceController.Get).Methods("GET")

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

	//equipment
	equipmentRepository := &repo.EquipmentRepository{DatabaseConnection: database}
	equipmentService := &service.EquipmentService{EquipmentRepo: equipmentRepository}
	equipmentController := &controller.EquipmentController{EquipmentService: equipmentService}

	keyPointRepository := &repo.KeyPointRepository{DatabaseConnection: database}
	keyPointService := &service.KeyPointService{KeyPointRepo: keyPointRepository}
	keyPointController := &controller.KeyPointController{KeyPointService: keyPointService}

	tourPreferenceRepository := &repo.TourPreferenceRepository{DatabaseConnection: database}
	tourPreferenceService := &service.TourPreferenceService{TourPreferenceRepo: tourPreferenceRepository}
	tourPreferenceController := &controller.TourPreferenceController{TourPreferenceService: tourPreferenceService}

	startServer(tourController, equipmentController, keyPointController, tourPreferenceController)
}
