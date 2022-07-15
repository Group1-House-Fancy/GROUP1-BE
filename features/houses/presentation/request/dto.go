package request

import "capstoneproject/features/houses"

type House struct {
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
	UserID       int     `json:"user_id" form:"user_id"`
}

func ToCore(req House) houses.Core {
	return houses.Core{
		Title:        req.Title,
		Price:        req.Price,
		Location:     req.Location,
		Longitude:    req.Longitude,
		Latitude:     req.Latitude,
		SurfaceArea:  req.SurfaceArea,
		BuildingArea: req.BuildingArea,
		Bathroom:     req.Bathroom,
		Bedroom:      req.Bedroom,
		Certificate:  req.Certificate,
		Description:  req.Description,
		Status:       req.Status,
		User: houses.User{
			ID: req.UserID,
		},
	}
}
