package controller

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TourPreferenceController struct {
	TourPreferenceService *service.TourPreferenceService
}

func (controller *TourPreferenceController) Create(writer http.ResponseWriter, req *http.Request) {
	var tourPreference model.TourPreference
	err := json.NewDecoder(req.Body).Decode(&tourPreference)

	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = tourPreference.Validate()
	if err != nil {
		println("Invalid data!")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = controller.TourPreferenceService.Create(&tourPreference)
	if err != nil {
		println("Error while creating a new tour preference")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	tourPreferenceJson, err := json.Marshal(tourPreference)
	if err != nil {
		println("Error while encoding tour preference to json")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Write(tourPreferenceJson)
}

func (controller *TourPreferenceController) Get(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["userId"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	tourPreference, err := controller.TourPreferenceService.Find(id)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	tourPreferenceJson, err := json.Marshal(tourPreference)
	if err != nil {
		println("Error while encoding tour preference to JSON")
	}
	writer.Write(tourPreferenceJson)
}
