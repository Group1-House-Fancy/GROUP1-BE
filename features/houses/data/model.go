package data

import (
	"capstoneproject/features/houses"

	"gorm.io/gorm"
)

type House struct {
	gorm.Model
	Title        string  `json:"title" form:"title"`
	Price        int     `json:"price" form:"price"`
	Location     string  `json:"location" form:"location"`
	Longitude    float64 `json:"longitude" form:"longitude"`
	Latitude     float64 `json:"latitude" form:"latitude"`
	SurfaceArea  int     `json:"surface_area" form:"surface_area"`
	BuildingArea int     `json:"building_area" form:"building_area"`
	Bathroom     int     `json:"bathroom" form:"bathroom"`
	Bedroom      int     `json:"bedroom" form:"bedroom"`
	Certificate  string  `json:"certificate" form:"certificate"`
	Description  string  `json:"description" form:"decsription"`
	Status       string  `json:"status" form:"status"`
	UserID       uint    `json:"user_id" form:"user_id"`
	User         User
	HouseImage   []HouseImage
}

type User struct {
	gorm.Model
	FullName    string `json:"full_name" form:"full_name"`
	Email       string `gorm:"unique" json:"email"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Address     string `json:"address" form:"address"`
	ImageURL    string `json:"image_url" form:"image_url"`
	House       []House
}

type HouseImage struct {
	gorm.Model
	HouseID  uint   `json:"house_id" form:"house_id"`
	ImageURL string `json:"image_url" form:"image_url"`
	House    House
}

func (data *House) toCore() houses.Core {
	return houses.Core{
		ID:           int(data.ID),
		Title:        data.Title,
		Price:        data.Price,
		Location:     data.Location,
		SurfaceArea:  data.SurfaceArea,
		BuildingArea: data.BuildingArea,
		Bathroom:     data.Bathroom,
		Bedroom:      data.Bedroom,
		Certificate:  data.Certificate,
		Description:  data.Description,
		Status:       data.Status,
		Longitude:    data.Longitude,
		Latitude:     data.Latitude,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
		User: houses.User{
			ID:          int(data.User.ID),
			FullName:    data.User.FullName,
			Email:       data.User.Email,
			PhoneNumber: data.User.PhoneNumber,
			Address:     data.User.Address,
			ImageURL:    data.User.ImageURL,
		},
		HouseImage: toCoreListHouseImages(data.HouseImage),
	}
}

func toCoreList(data []House) []houses.Core {
	result := []houses.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core houses.Core) House {
	return House{
		Title:        core.Title,
		Price:        core.Price,
		Location:     core.Location,
		Longitude:    core.Longitude,
		Latitude:     core.Latitude,
		SurfaceArea:  core.SurfaceArea,
		BuildingArea: core.BuildingArea,
		Bathroom:     core.Bathroom,
		Bedroom:      core.Bedroom,
		Certificate:  core.Certificate,
		Description:  core.Description,
		Status:       core.Status,
		UserID:       uint(core.User.ID),
	}
}

func (data *HouseImage) toCoreHouseImages() houses.HouseImage {
	return houses.HouseImage{
		ID:       int(data.ID),
		ImageURL: data.ImageURL,
		House: houses.Core{
			ID: int(data.HouseID),
		},
	}
}

func toCoreListHouseImages(data []HouseImage) []houses.HouseImage {
	result := []houses.HouseImage{}
	for key := range data {
		result = append(result, data[key].toCoreHouseImages())
	}
	return result
}
