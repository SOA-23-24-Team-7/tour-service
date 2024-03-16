package controller

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type KeyPointController struct {
	KeyPointService *service.KeyPointService
}

func (controller *KeyPointController) Create(writer http.ResponseWriter, req *http.Request) {
	var keyPoint model.KeyPoint
	err := json.NewDecoder(req.Body).Decode(&keyPoint)
	
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = keyPoint.Validate()
	if err != nil {
		println("Invalid data!")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = controller.KeyPointService.Create(&keyPoint)
	if err != nil {
		println("Error while creating a new key point")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	keyPointJson, err := json.Marshal(keyPoint)
	if err != nil {
    println("Error while encoding key point to json")
	}
    writer.WriteHeader(http.StatusInternalServerError)
	writer.Write(keyPointJson)
}

func (controller *KeyPointController) GetAll(writer http.ResponseWriter, req *http.Request){
	tourIdStr := mux.Vars(req)["tourId"]
	tourId, err := strconv.ParseInt(tourIdStr, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	keyPoints, err := controller.KeyPointService.FindAll(tourId)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	keyPointsJson,err := json.Marshal(keyPoints)
	if err != nil {
		println("Error while encoding key points to JSON")
		}
	writer.Write(keyPointsJson) 
}

func (controller *KeyPointController) Get(writer http.ResponseWriter, req *http.Request){
	idStr := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	keyPoint,err := controller.KeyPointService.Find(id)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	keyPointJson,err := json.Marshal(keyPoint)
	if err != nil {
		println("Error while encoding key point to JSON")
		}
	writer.Write(keyPointJson) 
}
