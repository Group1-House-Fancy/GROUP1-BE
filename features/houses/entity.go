package houses

import "time"

type Core struct {
	ID           int
	Title        string
	Price        int
	Location     string
	Longitude    float64
	Latitude     float64
	SurfaceArea  int
	BuildingArea int
	Bathroom     int
	Bedroom      int
	Certificate  string
	Description  string
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	User         User
}

type User struct {
	ID          int
	FullName    string
	Email       string
	PhoneNumber string
	Address     string
	ImageURL    string
}

type Business interface {
	GetAllHouse(limit, offset int) (data []Core, totalPage int, err error)
}

type Data interface {
	SelectAllHouse(limit, offset int) (data []Core, err error)
}