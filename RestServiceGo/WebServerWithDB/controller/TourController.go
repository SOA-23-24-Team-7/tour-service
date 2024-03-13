package controller

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
)

type TourController struct {
	TourService *service.TourService
}

func (controller *TourController) Create(writer http.ResponseWriter, req *http.Request) {
	var tour model.Tour
	err := json.NewDecoder(req.Body).Decode(&tour)
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
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tour)
}