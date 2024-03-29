package data

import (
	"capstoneproject/features/negotiations"

	"gorm.io/gorm"
)

type Negotiation struct {
	gorm.Model
	Nego    int    `json:"nego" form:"nego"`
	Status  string `json:"status" form:"status"`
	UserID  uint   `json:"user_id" form:"user_id"`
	HouseID uint   `json:"house_id" form:"house_id"`
	User    User
	House   House
}

type User struct {
	gorm.Model
	FullName    string `json:"full_name" form:"full_name"`
	ImageURL    string `json:"image_url" form:"image_url"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Email       string `gorm:"unique" json:"email"`
	Negotiation []Negotiation
}

type House struct {
	gorm.Model
	Title        string `json:"title" form:"title"`
	Price        int    `json:"price" form:"price"`
	Location     string `json:"location" form:"location"`
	SurfaceArea  int    `json:"surface_area" form:"surface_area"`
	BuildingArea int    `json:"building_area" form:"building_area"`
	Status       string `json:"status" form:"status"`
	Negotiation  []Negotiation
	HouseImage   []HouseImage
}

type HouseImage struct {
	gorm.Model
	HouseID  uint   `json:"house_id" form:"house_id"`
	ImageURL string `json:"image_url" form:"image_url"`
	House    House
}

func (data *Negotiation) toCore() negotiations.Core {
	return negotiations.Core{
		ID:        int(data.ID),
		Nego:      data.Nego,
		Status:    data.Status,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		User: negotiations.User{
			ID:          int(data.UserID),
			FullName:    data.User.FullName,
			ImageURL:    data.User.ImageURL,
			PhoneNumber: data.User.PhoneNumber,
			Email:       data.User.Email,
		},
		House: negotiations.House{
			ID:           int(data.HouseID),
			Title:        data.House.Title,
			Price:        data.House.Price,
			Location:     data.House.Location,
			SurfaceArea:  data.House.SurfaceArea,
			BuildingArea: data.House.BuildingArea,
			Status:       data.House.Status,
			HouseImage:   toCoreListHouseImages(data.House.HouseImage),
		},
	}
}

func toCoreList(data []Negotiation) []negotiations.Core {
	result := []negotiations.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core negotiations.Core) Negotiation {
	return Negotiation{
		Nego:    core.Nego,
		Status:  core.Status,
		UserID:  uint(core.User.ID),
		HouseID: uint(core.House.ID),
	}
}

func (data *HouseImage) toCoreHouseImages() negotiations.HouseImage {
	return negotiations.HouseImage{
		ID:       int(data.ID),
		ImageURL: data.ImageURL,
		House: negotiations.House{
			ID: int(data.HouseID),
		},
	}
}

func toCoreListHouseImages(data []HouseImage) []negotiations.HouseImage {
	result := []negotiations.HouseImage{}
	for key := range data {
		result = append(result, data[key].toCoreHouseImages())
	}
	return result
}
