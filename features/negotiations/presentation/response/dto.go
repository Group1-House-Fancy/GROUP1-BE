package response

import (
	"capstoneproject/features/negotiations"
	"time"
)

type History struct {
	ID        int       `json:"id" form:"id"`
	Nego      int       `json:"nego" form:"nego"`
	Status    string    `json:"status" form:"status"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	House     House     `json:"house" form:"house"`
}

type Negotiator struct {
	ID   int  `json:"id" form:"id"`
	Nego int  `json:"nego" form:"nego"`
	User User `json:"user" form:"user"`
}

type House struct {
	ID           int    `json:"id" form:"id"`
	Title        string `json:"title" form:"title"`
	Price        int    `json:"price" form:"price"`
	Location     string `json:"location" form:"location"`
	SurfaceArea  int    `json:"surface_area" form:"surface_area"`
	BuildingArea int    `json:"building_area" form:"building_area"`
}

type User struct {
	ID          int    `json:"id" form:"id"`
	FullName    string `json:"full_name" form:"full_name"`
	ImageURL    string `json:"image_url" form:"image_url"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
}

func FromCoreHistory(data negotiations.Core) History {
	return History{
		ID:        data.ID,
		Nego:      data.Nego,
		Status:    data.Status,
		CreatedAt: data.CreatedAt,
		House: House{
			ID:           data.House.ID,
			Title:        data.House.Title,
			Price:        data.House.Price,
			Location:     data.House.Location,
			SurfaceArea:  data.House.SurfaceArea,
			BuildingArea: data.House.BuildingArea,
		},
	}
}

func FromCoreHistoryList(data []negotiations.Core) []History {
	result := []History{}
	for key := range data {
		result = append(result, FromCoreHistory(data[key]))
	}
	return result
}

func FromCoreNegotiator(data negotiations.Core) Negotiator {
	return Negotiator{
		ID:   data.ID,
		Nego: data.Nego,
		User: User{
			ID:          data.User.ID,
			FullName:    data.User.FullName,
			ImageURL:    data.User.ImageURL,
			PhoneNumber: data.User.PhoneNumber,
		},
	}
}

func FromCoreNegotiatorList(data []negotiations.Core) []Negotiator {
	result := []Negotiator{}
	for key := range data {
		result = append(result, FromCoreNegotiator(data[key]))
	}
	return result
}
