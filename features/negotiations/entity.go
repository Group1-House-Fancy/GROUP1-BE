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
	ID       int
	FullName string
	ImageURL string
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
}

type Data interface {
	SelectNegotiationsByIdUser(idUser, limit, offset int) (data []Core, err error)
}
