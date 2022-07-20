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
	HouseImage   []HouseImage
}

type User struct {
	ID          int
	FullName    string
	Email       string
	PhoneNumber string
	Address     string
	ImageURL    string
}

type HouseImage struct {
	ID       int
	ImageURL string
	House    Core
}

type Business interface {
	GetAllHouse(limit, offset int) (data []Core, totalPage int, err error)
	PostNewHouse(data Core) (idHouse int, row int, err error)
	GetHouseDetail(idHouse int) (data Core, err error)
	GetMyListHouse(idUser, limit, offset int) (data []Core, totalPage int, err error)
	PutHouse(idHouse int, data Core) (row int, err error)
	DeleteHouse(idHouse int) (row int, err error)
	GetSearchHouse(keywords string, limit, offset int) (data []Core, totalPage int, err error)
}

type Data interface {
	SelectAllHouse(limit, offset int) (data []Core, err error)
	InsertNewHouse(data Core) (idHouse int, row int, err error)
	SelectHouseByIdHouse(idHouse int) (data Core, err error)
	SelectHouseByIdUser(idUser, limit, offset int) (data []Core, err error)
	UpdateHouse(idHouse int, data Core) (row int, err error)
	DeleteHouse(idHouse int) (row int, err error)
	SelectSearchHouse(keywords string, limit, offset int) (data []Core, err error)
}
