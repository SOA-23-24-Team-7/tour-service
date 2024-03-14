package controller

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type TourController struct {
	TourService *service.TourService
}

func (controller *TourController) Create(writer http.ResponseWriter, req *http.Request) {
	var tour model.Tour
	err := json.NewDecoder(req.Body).Decode(&tour)
	fmt.Println(tour)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = controller.TourService.Create(&tour)
	if err != nil {
		println("Error while creating a new student")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	tourJSON, err := json.Marshal(tour)
	if err != nil {
    println("Error while encoding tour to JSON")
	}
    writer.WriteHeader(http.StatusInternalServerError)
	writer.Write(tourJSON)
    
//ovdje je bio jedan return 
}