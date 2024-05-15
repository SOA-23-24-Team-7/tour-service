package server

// import (
// 	"context"
// 	"database-example/model"
// 	"database-example/service"
// 	"encoding/json"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"
// )

// type TourServer struct {
// 	UnimplementedTourServerServer
// 	TourService *service.TourService
// }

// func (server *TourServer) AddEquipment(writer http.ResponseWriter, req *http.Request) {
// 	tourIdStr := mux.Vars(req)["tourId"]
// 	tourId, err := strconv.ParseInt(tourIdStr, 10, 64)
// 	if err != nil {
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	equipmentIdStr := mux.Vars(req)["equipmentId"]
// 	equipmentId, err := strconv.ParseInt(equipmentIdStr, 10, 64)
// 	if err != nil {
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	errorr := server.TourService.AddEquipment(tourId, equipmentId)
// 	if errorr != nil {
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	writer.WriteHeader(http.StatusOK)
// 	return

// }

// func (server *TourServer) GetEquipment(writer http.ResponseWriter, req *http.Request) {
// 	tourIdStr := mux.Vars(req)["tourId"]
// 	tourId, err := strconv.ParseInt(tourIdStr, 10, 64)
// 	if err != nil {
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	equipment, errorr := server.TourService.GetEquipment(tourId)
// 	if errorr != nil {
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	if equipment == nil {
// 		equipment = make([]model.Equipment, 0)
// 	}
// 	writer.WriteHeader(http.StatusOK)
// 	json.NewEncoder(writer).Encode(equipment)

// }

// func (server *TourServer) DeleteEquipment(writer http.ResponseWriter, req *http.Request) {
// 	tourIdStr := mux.Vars(req)["tourId"]
// 	tourId, err := strconv.ParseInt(tourIdStr, 10, 64)
// 	if err != nil {
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	equipmentIdStr := mux.Vars(req)["equipmentId"]
// 	equipmentId, err := strconv.ParseInt(equipmentIdStr, 10, 64)
// 	if err != nil {
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	errorr := server.TourService.DeleteEquipment(tourId, equipmentId)
// 	if errorr != nil {
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	writer.WriteHeader(http.StatusOK)

// }
