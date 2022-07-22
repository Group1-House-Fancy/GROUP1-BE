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
	GetAllPortfolio(idCtr, limit, offset int) (data []Core, totalPage int, err error)
	GetPortfolio(idPrtf int) (data Core, err error)
	PutPortfolio(idPrtf int, data Core) (row int, err error)
	DeletePortfolio(idPrtf int) (row int, err error)
}

type Data interface {
	InsertPortfolio(data Core) (row int, idPrt int, err error)
	SelectAllPortfolio(idCtr, limit, offset int) (data []Core, err error)
	SelectPortfolio(idPrtf int) (data Core, err error)
	UpdatePortfolio(idPrtf int, data Core) (row int, err error)
	DeletePortfolio(idPrtf int) (row int, err error)
	CountPortfolioData(idCtr int) (count int, err error)
}
