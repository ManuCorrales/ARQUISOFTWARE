package services

import (
	Amenitieclient "mvc-go/clients/amenitie"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type amenitieService struct{}


type amenitieServiceInterface interface {
	GetAmenities() (dto.AmenitiesDto, e.ApiError)
	GetAmenitieById(id int) (dto.AmenitieDto, e.ApiError)
	InsertAmenitie(amenitieDto dto.AmenitieDto) (dto.AmenitieDto, e.ApiError)
	DeleteAmenitieById(amenitieId int) e.ApiError
}

var (
	AmenitieService amenitieServiceInterface
)

func init() {
	AmenitieService = &amenitieService{}
}

func (a *amenitieService) GetAmenities() (dto.AmenitiesDto, e.ApiError) {

	var amenities model.Amenities = Amenitieclient.GetAmenities()
	amenitiesList := make([]dto.AmenitieDto, 0)

	for _, amenitie := range amenities {
		var amenitieDto dto.AmenitieDto

		amenitieDto.Description = amenitie.Description
		amenitieDto.Id = amenitie.Id
		amenitieDto.Name = amenitie.Name

		amenitiesList = append(amenitiesList, amenitieDto)
	}

	return dto.AmenitiesDto{
		Amenities: amenitiesList,
	}, nil
}

func (a *amenitieService) GetAmenitieById(id int) (dto.AmenitieDto, e.ApiError) {

	var amenitie model.Amenitie = Amenitieclient.GetAmenitieById(id)
	var amenitieDto dto.AmenitieDto

	if amenitie.Id == 0 {
		return amenitieDto, e.NewBadRequestApiError("no se ha encontrado la reserva")
	}

	amenitieDto.Id = amenitie.Id
	amenitieDto.Name = amenitie.Name
	amenitieDto.Description = amenitie.Description

	return amenitieDto, nil
}

func (a *amenitieService) InsertAmenitie(amenitieDto dto.AmenitieDto) (dto.AmenitieDto, e.ApiError) {

	var amenitie model.Amenitie

	amenitie.Description = amenitieDto.Description
	amenitie.Name = amenitieDto.Name

	amenitie = Amenitieclient.InsertAmenitie(amenitie)

	amenitieDto.Id = amenitie.Id

	return amenitieDto, nil
}

func (a *amenitieService) DeleteAmenitieById(amenitieId int) e.ApiError {
	_, err := a.GetAmenitieById(amenitieId)

	if err != nil {
		return err
	}

	err = Amenitieclient.DeleteAmenitieById(amenitieId)

	if err != nil {
		return e.NewInternalServerApiError("Failed to delete amenitie", err)
	}

	return nil
}