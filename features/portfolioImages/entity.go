package portfolioImages

import "time"

type Core struct {
	ID        int
	ImageURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Portfolio Portfolio
}

type Portfolio struct {
	ID int
}

type Business interface {
	PostNewPortfolioImage(data Core) (row int, err error)
	DeleteImage(idImage int) (row int, err error)
}

type Data interface {
	InsertNewImage(data Core) (row int, err error)
	DeleteImage(idImage int) (row int, err error)
}
