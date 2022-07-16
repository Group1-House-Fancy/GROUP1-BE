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
}

type Data interface {
	InsertNewImages(data Core) (row int, err error)
}
