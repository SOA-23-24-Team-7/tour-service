package main

import (
	"database-example/model"
	"database-example/repo"
	"database-example/server"
	"database-example/service"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := "user=postgres password=super dbname=soa-tours host=tours-database port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	database.AutoMigrate(&model.Tour{})
	database.AutoMigrate(&model.Equipment{})
	database.AutoMigrate(&model.KeyPoint{})
	database.AutoMigrate(&model.TourPreference{})
	database.AutoMigrate(&model.Facility{})

	return database
}

func startServer(tourService *service.TourService,
	equipmentService *service.EquipmentService,
	keyPointService *service.KeyPointService,
	tourPreferenceService *service.TourPreferenceService,
	facilityService *service.FacilityService) {

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	server.RegisterTourMicroserviceServer(grpcServer, &server.TourMicroservice{
		TourPreferenceService: tourPreferenceService,
		FacilityService:       facilityService,
		KeyPointService:       keyPointService,
		TourService:           tourService,
		EquipmentService:      equipmentService,
	})

	listener, err := net.Listen("tcp", ":8087")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("gRPC server listening on port :8087")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

func main() {

	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	equipmentRepository := &repo.EquipmentRepository{DatabaseConnection: database}
	equipmentService := &service.EquipmentService{EquipmentRepo: equipmentRepository}

	tourRepository := &repo.TourRepository{DatabaseConnection: database}
	tourService := &service.TourService{TourRepo: tourRepository,
		EquipmentRepo: equipmentRepository}

	keyPointRepository := &repo.KeyPointRepository{DatabaseConnection: database}
	keyPointService := &service.KeyPointService{KeyPointRepo: keyPointRepository}

	tourPreferenceRepository := &repo.TourPreferenceRepository{DatabaseConnection: database}
	tourPreferenceService := &service.TourPreferenceService{TourPreferenceRepo: tourPreferenceRepository}

	facilityRepository := &repo.FacilityRepository{DatabaseConnection: database}
	facilityService := &service.FacilityService{FacilityRepo: facilityRepository}

	startServer(tourService, equipmentService, keyPointService, tourPreferenceService, facilityService)
}
