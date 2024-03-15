package controller

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type EquipmentController struct {
	EquipmentService *service.EquipmentService
}

func (controller *EquipmentController) Create(writer http.ResponseWriter, req *http.Request) {
	var equipment model.Equipment
	err := json.NewDecoder(req.Body).Decode(&equipment)
	
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = equipment.Validate()
	if err != nil {
		println("Invalid data!")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = controller.EquipmentService.Create(&equipment)
	if err != nil {
		println("Error while creating a new student")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	tourJSON, err := json.Marshal(equipment)
	if err != nil {
    println("Error while encoding tour to JSON")
	}
    writer.WriteHeader(http.StatusInternalServerError)
	writer.Write(tourJSON)
    

}

func (controller *EquipmentController) GetAll(writer http.ResponseWriter, req *http.Request){
	
	equipment,err := controller.EquipmentService.FindAll()
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	equipmentJson,err := json.Marshal(equipment)
	if err != nil {
		println("Error while encoding tour to JSON")
		}
	writer.Write(equipmentJson) 
}

func (controller *EquipmentController) Get(writer http.ResponseWriter, req *http.Request){
	idStr := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	equipment,err := controller.EquipmentService.Find(id)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	equipmentJson,err := json.Marshal(equipment)
	if err != nil {
		println("Error while encoding tour to JSON")
		}
	writer.Write(equipmentJson) 

}