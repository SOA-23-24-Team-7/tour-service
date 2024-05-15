package server

import (
	"context"
	"database-example/model"
	"database-example/service"
)

type TourMicroservice struct {
	UnimplementedTourMicroserviceServer
	TourPreferenceService *service.TourPreferenceService
	FacilityService       *service.FacilityService
	KeyPointService       *service.KeyPointService
	TourService           *service.TourService
	EquipmentService      *service.EquipmentService
}

func (server *TourMicroservice) CreatePreference(ctx context.Context, req *TourPreferenceCreationRequest) (*TourPreferenceResponse, error) {
	tourPreference := &model.TourPreference{
		UserId:          req.UserId,
		DifficultyLevel: req.DifficultyLevel,
		WalkingRating:   req.WalkingRating,
		CyclingRating:   req.CyclingRating,
		CarRating:       req.CarRating,
		BoatRating:      req.BoatRating,
		SelectedTags:    req.SelectedTags,
	}

	if err := tourPreference.Validate(); err != nil {
		return nil, err
	}

	if err := server.TourPreferenceService.Create(tourPreference); err != nil {
		return nil, err
	}

	response := &TourPreferenceResponse{
		Id:              tourPreference.Id,
		UserId:          tourPreference.UserId,
		DifficultyLevel: tourPreference.DifficultyLevel,
		WalkingRating:   tourPreference.WalkingRating,
		CyclingRating:   tourPreference.CyclingRating,
		CarRating:       tourPreference.CarRating,
		BoatRating:      tourPreference.BoatRating,
		SelectedTags:    tourPreference.SelectedTags,
	}

	return response, nil
}

func (server *TourMicroservice) GetPreference(ctx context.Context, req *TourPreferenceIdRequest) (*TourPreferenceResponse, error) {
	tourPreference, err := server.TourPreferenceService.Find(req.UserId)
	if err != nil {
		return nil, err
	}

	response := &TourPreferenceResponse{
		Id:              tourPreference.Id,
		UserId:          tourPreference.UserId,
		DifficultyLevel: tourPreference.DifficultyLevel,
		WalkingRating:   tourPreference.WalkingRating,
		CyclingRating:   tourPreference.CyclingRating,
		CarRating:       tourPreference.CarRating,
		BoatRating:      tourPreference.BoatRating,
		SelectedTags:    tourPreference.SelectedTags,
	}

	return response, nil
}

func (server *TourMicroservice) CreateFacility(ctx context.Context, req *FacilityCreationRequest) (*FacilityResponse, error) {
	facility := &model.Facility{
		Name:        req.Name,
		Description: req.Description,
		ImagePath:   req.ImagePath,
		AuthorId:    req.AuthorId,
		Category:    model.FacilityCategory(req.Category),
		Longitude:   req.Longitude,
		Latitude:    req.Latitude,
	}

	if err := server.FacilityService.Create(facility); err != nil {
		return nil, err
	}

	response := &FacilityResponse{
		Id:          facility.Id,
		Name:        facility.Name,
		Description: facility.Description,
		ImagePath:   facility.ImagePath,
		AuthorId:    facility.AuthorId,
		Category:    int32(facility.Category),
		Longitude:   facility.Longitude,
		Latitude:    facility.Latitude,
	}

	return response, nil
}

func (server *TourMicroservice) GetAllFacilities(ctx context.Context, req *FacilitiesIdRequest) (*FacilityListResponse, error) {
	facilities, err := server.FacilityService.FindAll(req.AuthorId)
	if err != nil {
		return nil, err
	}

	var response FacilityListResponse
	response.Facilities = []*FacilityResponse{}
	for _, f := range facilities {
		response.Facilities = append(response.Facilities, &FacilityResponse{
			Id:          f.Id,
			Name:        f.Name,
			Description: f.Description,
			ImagePath:   f.ImagePath,
			AuthorId:    f.AuthorId,
			Category:    int32(f.Category),
			Longitude:   f.Longitude,
			Latitude:    f.Latitude,
		})
	}

	return &response, nil
}

func (server *TourMicroservice) CreateKeyPoint(ctx context.Context, req *KeyPointCreationRequest) (*KeyPointResponse, error) {
	var keyPointSecret *model.KeyPointSecret
	if req.Secret != nil {
		keyPointSecret = &model.KeyPointSecret{
			Images:      req.Secret.Images,
			Description: req.Secret.Description,
		}
	}

	keyPoint := &model.KeyPoint{
		TourId:          req.TourId,
		Name:            req.Name,
		Description:     req.Description,
		Longitude:       req.Longitude,
		Latitude:        req.Latitude,
		LocationAddress: req.LocationAddress,
		ImagePath:       req.ImagePath,
		Order:           req.Order,
		HaveSecret:      keyPointSecret != nil,
		KeyPointSecret:  *keyPointSecret,
	}

	if err := keyPoint.Validate(); err != nil {
		return nil, err
	}

	if err := server.KeyPointService.Create(keyPoint); err != nil {
		return nil, err
	}

	response := &KeyPointResponse{
		Id:              keyPoint.Id,
		TourId:          keyPoint.TourId,
		Name:            keyPoint.Name,
		Description:     keyPoint.Description,
		Longitude:       keyPoint.Longitude,
		Latitude:        keyPoint.Latitude,
		LocationAddress: keyPoint.LocationAddress,
		ImagePath:       keyPoint.ImagePath,
		Order:           keyPoint.Order,
		HaveSecret:      keyPoint.HaveSecret,
		Secret: &KeyPointSecretResponse{
			Images:      keyPoint.KeyPointSecret.Images,
			Description: keyPoint.KeyPointSecret.Description,
		},
	}

	return response, nil
}

func (server *TourMicroservice) GetAllKeyPoints(ctx context.Context, req *KeyPointsIdRequest) (*KeyPointListResponse, error) {
	keyPoints, err := server.KeyPointService.FindAll(req.TourId)
	if err != nil {
		return nil, err
	}

	var response KeyPointListResponse
	response.KeyPoints = []*KeyPointResponse{}
	for _, kp := range keyPoints {
		var keyPointSecretResponse *KeyPointSecretResponse
		if kp.HaveSecret {
			keyPointSecretResponse = &KeyPointSecretResponse{
				Images:      kp.KeyPointSecret.Images,
				Description: kp.KeyPointSecret.Description,
			}
		}

		response.KeyPoints = append(response.KeyPoints, &KeyPointResponse{
			Id:              kp.Id,
			TourId:          kp.TourId,
			Name:            kp.Name,
			Description:     kp.Description,
			Longitude:       kp.Longitude,
			Latitude:        kp.Latitude,
			LocationAddress: kp.LocationAddress,
			ImagePath:       kp.ImagePath,
			Order:           kp.Order,
			HaveSecret:      kp.HaveSecret,
			Secret:          keyPointSecretResponse,
		})
	}

	return &response, nil
}

func (server *TourMicroservice) GetKeyPoint(ctx context.Context, req *KeyPointIdRequest) (*KeyPointResponse, error) {
	keyPoint, err := server.KeyPointService.Find(req.Id)
	if err != nil {
		return nil, err
	}

	var keyPointSecretResponse *KeyPointSecretResponse
	if keyPoint.HaveSecret {
		keyPointSecretResponse = &KeyPointSecretResponse{
			Images:      keyPoint.KeyPointSecret.Images,
			Description: keyPoint.KeyPointSecret.Description,
		}
	}

	response := &KeyPointResponse{
		Id:              keyPoint.Id,
		TourId:          keyPoint.TourId,
		Name:            keyPoint.Name,
		Description:     keyPoint.Description,
		Longitude:       keyPoint.Longitude,
		Latitude:        keyPoint.Latitude,
		LocationAddress: keyPoint.LocationAddress,
		ImagePath:       keyPoint.ImagePath,
		Order:           keyPoint.Order,
		HaveSecret:      keyPoint.HaveSecret,
		Secret:          keyPointSecretResponse,
	}

	return response, nil
}

func (server *TourMicroservice) CreateTour(ctx context.Context, req *TourCreationRequest) (*TourResponse, error) {
	tour := &model.Tour{
		Id:          req.Id,
		AuthorId:    req.AuthorId,
		Name:        req.Name,
		Description: req.Description,
		Difficulty:  req.Difficulty,
		Tags:        req.Tags,
		Status:      model.TourStatus(req.Status),
		Price:       req.Price,
		IsDeleted:   req.IsDeleted,
		Distance:    req.Distance,
		Category:    model.TourCategory(req.Category),
	}

	if err := server.TourService.Create(tour); err != nil {
		return nil, err
	}

	response := &TourResponse{
		Id:          tour.Id,
		AuthorId:    tour.AuthorId,
		Name:        tour.Name,
		Description: tour.Description,
		Difficulty:  tour.Difficulty,
		Tags:        tour.Tags,
		Status:      int32(tour.Status),
		Price:       tour.Price,
		IsDeleted:   tour.IsDeleted,
		Distance:    tour.Distance,
		Category:    int32(tour.Category),
	}

	return response, nil
}

func (server *TourMicroservice) GetAllTours(ctx context.Context, req *ToursIdRequest) (*TourListResponse, error) {
	tours, err := server.TourService.FindAll(req.AuthorId)
	if err != nil {
		return nil, err
	}

	var response TourListResponse
	response.Tours = []*TourResponse{}
	for _, t := range tours {
		response.Tours = append(response.Tours, &TourResponse{
			Id:          t.Id,
			AuthorId:    t.AuthorId,
			Name:        t.Name,
			Description: t.Description,
			Difficulty:  int32(t.Difficulty),
			Tags:        t.Tags,
			Status:      int32(t.Status),
			Price:       t.Price,
			IsDeleted:   t.IsDeleted,
			Distance:    t.Distance,
			Category:    int32(t.Category),
		})
	}

	return &response, nil
}

func (server *TourMicroservice) GetTour(ctx context.Context, req *TourIdRequest) (*TourResponse, error) {
	tour, err := server.TourService.Find(req.Id)
	if err != nil {
		return nil, err
	}

	response := &TourResponse{
		Id:          tour.Id,
		AuthorId:    tour.AuthorId,
		Name:        tour.Name,
		Description: tour.Description,
		Difficulty:  int32(tour.Difficulty),
		Tags:        tour.Tags,
		Status:      int32(tour.Status),
		Price:       tour.Price,
		IsDeleted:   tour.IsDeleted,
		Distance:    tour.Distance,
		Category:    int32(tour.Category),
	}

	return response, nil
}

func (server *TourMicroservice) AddTourEquipment(ctx context.Context, req *TourEquipmentCreationRequest) (*Empty, error) {
	err := server.TourService.AddEquipment(req.TourId, req.EquipmentId)
	if err != nil {
		return nil, err
	}

	return &Empty{}, nil
}

func (server *TourMicroservice) GetTourEquipment(ctx context.Context, req *TourEquipmentListIdRequest) (*EquipmentListResponse, error) {
	equipment, err := server.TourService.GetEquipment(req.TourId)
	if err != nil {
		return nil, err
	}

	var response EquipmentListResponse
	response.Equipment = []*EquipmentResponse{}
	for _, e := range equipment {
		response.Equipment = append(response.Equipment, &EquipmentResponse{
			Id:          e.Id,
			Name:        e.Name,
			Description: e.Description,
		})
	}

	return &response, nil
}

func (server *TourMicroservice) DeleteTourEquipment(ctx context.Context, req *TourEquipmentDeletionRequest) (*Empty, error) {
	err := server.TourService.DeleteEquipment(req.TourId, req.EquipmentId)
	if err != nil {
		return nil, err
	}

	return &Empty{}, nil
}

func (server *TourMicroservice) CreateEquipment(ctx context.Context, req *EquipmentCreationRequest) (*EquipmentResponse, error) {
	equipment := model.Equipment{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := server.EquipmentService.Create(&equipment); err != nil {
		return nil, err
	}

	response := &EquipmentResponse{
		Id:          equipment.Id,
		Name:        equipment.Name,
		Description: equipment.Description,
	}

	return response, nil
}

func (server *TourMicroservice) GetAllEquipment(ctx context.Context, req *Empty) (*EquipmentListResponse, error) {
	equipment, err := server.EquipmentService.FindAll()
	if err != nil {
		return nil, err
	}

	var response EquipmentListResponse
	for _, e := range equipment {
		response.Equipment = append(response.Equipment, &EquipmentResponse{
			Id:          e.Id,
			Name:        e.Name,
			Description: e.Description,
		})
	}

	return &response, nil
}

func (server *TourMicroservice) GetEquipment(ctx context.Context, req *EquipmentIdRequest) (*EquipmentResponse, error) {
	equipment, err := server.EquipmentService.Find(req.Id)
	if err != nil {
		return nil, err
	}

	if equipment == nil {
		return nil, err
	}

	response := &EquipmentResponse{
		Id:          equipment.Id,
		Name:        equipment.Name,
		Description: equipment.Description,
	}

	return response, nil
}
