package controller

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type FacilityController struct {
	FacilityService *service.FacilityService
}

func (controller *FacilityController) Create(writer http.ResponseWriter, req *http.Request) {
	var facility model.Facility
	err := json.NewDecoder(req.Body).Decode(&facility)

	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = facility.Validate()
	if err != nil {
		println("Invalid data!")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = controller.FacilityService.Create(&facility)
	if err != nil {
		println("Error while creating a new facility")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	facilityJson, err := json.Marshal(facility)
	if err != nil {
		println("Error while encoding facility to json")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Write(facilityJson)
}

func (controller *FacilityController) GetAll(writer http.ResponseWriter, req *http.Request) {
	authorIdStr := mux.Vars(req)["userId"]
	authorId, err := strconv.ParseInt(authorIdStr, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	facilities, err := controller.FacilityService.FindAll(authorId)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	facilitiesJson, err := json.Marshal(facilities)
	if err != nil {
		println("Error while encoding facilities to JSON")
	}
	writer.Write(facilitiesJson)
}
