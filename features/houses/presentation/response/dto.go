package response

import (
	"capstoneproject/features/houses"
	"time"
)

type House struct {
	ID           int       `json:"id" form:"id"`
	Title        string    `json:"title" form:"title"`
	Price        int       `json:"price" form:"price"`
	Location     string    `json:"location" form:"location"`
	Longitude    float64   `json:"longitude" form:"longitude"`
	Latitude     float64   `json:"latitude" form:"latitude"`
	SurfaceArea  int       `json:"surface_area" form:"surface_area"`
	BuildingArea int       `json:"building_area" form:"building_area"`
	Bathroom     int       `json:"bathroom" form:"bathroom"`
	Bedroom      int       `json:"bedroom" form:"bedroom"`
	Certificate  string    `json:"certificate" form:"certificate"`
	Description  string    `json:"description" form:"decsription"`
	Status       string    `json:"status" form:"status"`
	CreatedAt    time.Time `json:"created_at" form:"created_at"`
	User         User      `json:"user" form:"user"`
}

type User struct {
	ID          int    `json:"id" form:"id"`
	FullName    string `json:"full_name" form:"full_name"`
	Email       string `json:"email" form:"email"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Address     string `json:"address" form:"address"`
	ImageURL    string `json:"image_url" form:"image_url"`
}

func FromCore(data houses.Core) House {
	return House{
		ID:           data.ID,
		Title:        data.Title,
		Price:        data.Price,
		Location:     data.Location,
		Longitude:    data.Longitude,
		Latitude:     data.Latitude,
		SurfaceArea:  data.SurfaceArea,
		BuildingArea: data.BuildingArea,
		Bathroom:     data.Bathroom,
		Bedroom:      data.Bedroom,
		Certificate:  data.Certificate,
		Description:  data.Description,
		Status:       data.Status,
		CreatedAt:    data.CreatedAt,
		User: User{
			ID:          data.User.ID,
			FullName:    data.User.FullName,
			Email:       data.User.Email,
			Address:     data.User.Address,
			PhoneNumber: data.User.PhoneNumber,
			ImageURL:    data.User.ImageURL,
		},
	}
}
func FromCoreList(data []houses.Core) []House {
	result := []House{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
