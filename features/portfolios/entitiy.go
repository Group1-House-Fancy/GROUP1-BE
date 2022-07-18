package portfolios

import "time"

type Core struct {
	ID             int
	ClientName     string
	Location       string
	FinishDate     string
	Longitude      float64
	Latitude       float64
	Price          int
	Description    string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Contractor     Contractor
	PortfolioImage []PortfolioImage
}

type Contractor struct {
	ID int
}

type PortfolioImage struct {
	ID        int
	ImageURL  string
	Portfolio Core
}

type Business interface {
	PostPortfolio(data Core) (row int, idPrt int, err error)
}

type Data interface {
	InsertPortfolio(data Core) (row int, idPrt int, err error)
}
