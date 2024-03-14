package controller

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	err = tour.Validate()
	if err != nil {
		println("Invalid data!")
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
    

}

func (controller *TourController) GetAll(writer http.ResponseWriter, req *http.Request){
	idStr := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	tours,err := controller.TourService.FindAll(id)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	toursJson,err := json.Marshal(tours)
	if err != nil {
		println("Error while encoding tour to JSON")
		}
	writer.Write(toursJson) 
}

func (controller *TourController) Get(writer http.ResponseWriter, req *http.Request){
	idStr := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	tour,err := controller.TourService.Find(id)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	toursJson,err := json.Marshal(tour)
	if err != nil {
		println("Error while encoding tour to JSON")
		}
	writer.Write(toursJson) 

}
