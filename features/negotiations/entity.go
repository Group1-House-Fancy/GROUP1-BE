package negotiations

import "time"

type Core struct {
	ID        int
	Nego      int
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
	House     House
}

type User struct {
	ID          int
	FullName    string
	ImageURL    string
	PhoneNumber string
}

type House struct {
	ID           int
	Title        string
	Price        int
	Location     string
	SurfaceArea  int
	BuildingArea int
	Status       string
}

type Business interface {
	GetHistoryUser(idUser, limit, offset int) (data []Core, totalPage int, err error)
	GetHouseNegotiators(idHouse, limit, offset int) (data []Core, totalPage int, err error)
	PostNewNegotiation(data Core) (row int, err error)
	UpdateStatus(idNegotiation int, status string) (row int, err error)
}

type Data interface {
	SelectNegotiationsByIdUser(idUser, limit, offset int) (data []Core, err error)
	SelectNegotiationsByIdHouse(idHouse, limit, offset int) (data []Core, err error)
	SelectNegotiation(idNegotiation int) (data Core)
	InsertNewNegotiation(data Core) (row int, err error)
	CheckAlreadyNegotiation(idUser, idHouse int) (row int, err error)
	UpdateHouseStatus(idHouse int, status string) (row int, err error)
	UpdateNegotiation(idNegotiation int, status string) (row int, err error)
	CheckNegotiator(idHouse int) (cond bool)
}
