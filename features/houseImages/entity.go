package houseimages

import "time"

type Core struct {
	ID        int
	ImageURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
	House     House
}

type House struct {
	ID int
}

type Business interface {
	PostNewHouseImage(data Core) (row int, err error)
	DeleteImage(idImage int) (row int, err error)
}

type Data interface {
	InsertNewImage(data Core) (row int, err error)
	DeleteImage(idImage int) (row int, err error)
}
