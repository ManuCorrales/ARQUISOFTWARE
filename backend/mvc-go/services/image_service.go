package services


import (
	"io"
	"mime/multipart"
	"os"

	Imageclient "mvc-go/clients/image"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type imageService struct{}

type imageServiceInterface interface {
	GetImageById(id int) (dto.ImageDto, e.ApiError)
	GetImages() (dto.ImagesDto, e.ApiError)
	ImageInsert(hotelID int, URL string) (dto.ImageDto, e.ApiError)
	GetImagesByHotelId(hotelID int) (dto.ImagesDto, e.ApiError)
	DeleteImageById(id int) e.ApiError
}

var (
	ImageService imageServiceInterface
)

func init() {
	ImageService = &imageService{}
}

func (i *imageService) GetImageById(id int) (dto.ImageDto, e.ApiError) {
	image := Imageclient.GetImageById(id)

	if image.Id == 0 {
		return dto.ImageDto{}, e.NewNotFoundApiError("Image not found")
	}

	imageDto := dto.ImageDto{
		Id:  image.Id,
		Url: image.Url,
	}

	return imageDto, nil
}

func (i *imageService) GetImages() (dto.ImagesDto, e.ApiError) {
	images := Imageclient.GetImages()
	imageDtos := make([]dto.ImageDto, len(images))

	for i, image := range images {
		imageDto := dto.ImageDto{
			Id:      image.Id,
			Url:     image.Url,
			HotelId: image.HotelId,
		}
		imageDtos[i] = imageDto
	}

	return dto.ImagesDto{
		Images: imageDtos,
	}, nil
}

func (i *imageService) ImageInsert(hotelID int, URL string) (dto.ImageDto, e.ApiError) {

	var imageDto dto.ImageDto
	var image model.Image

	image.HotelId = hotelID
	image.Url = URL


	image = Imageclient.ImageInsert(image)

	if image.Id == 0 {
		return imageDto, e.NewBadRequestApiError("Error al insertar imagen")
	}

	imageDto.HotelId = image.HotelId
	imageDto.Id = image.Id
	imageDto.Url = image.Url

	return imageDto, nil

}

func saveFile(imageFile *multipart.FileHeader, filePath string) error {
	// Abrir el archivo cargado
	file, err := imageFile.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	// Crear el archivo destino en el sistema de archivos
	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copiar el contenido del archivo cargado al archivo destino
	_, err = io.Copy(dst, file)
	if err != nil {
		return err
	}

	return nil
}

func (i *imageService) GetImagesByHotelId(hotelID int) (dto.ImagesDto, e.ApiError) {
	images := Imageclient.GetImagesByHotelId(hotelID)
	imageDtos := make([]dto.ImageDto, len(images))

	for i, image := range images {
		imageDto := dto.ImageDto{
			Id:   image.Id,
			Url:  image.Url,
		}
		imageDtos[i] = imageDto
	}

	return dto.ImagesDto{
		Images: imageDtos,
	}, nil
}

func (i *imageService) DeleteImageById(id int) e.ApiError {
	// Verificar si la imagen existe
	_, err := i.GetImageById(id)
	if err != nil {
		return err
	}

	err = Imageclient.DeleteImageById(id)
	if err != nil {
		return e.NewInternalServerApiError("Failed to delete image", err)
	}

	return nil 
}



