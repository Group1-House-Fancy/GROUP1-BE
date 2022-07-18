package request

import "capstoneproject/features/portfolios"

type Portfolio struct {
	ClientName   string  `json:"client_name" form:"client_name"`
	Location     string  `json:"location" form:"location"`
	FinishDate   string  `json:"finish_date" form:"finish_date"`
	Longitude    float64 `json:"longitude" form:"longitude"`
	Latitude     float64 `json:"latitude" form:"latitude"`
	Price        int     `json:"price" form:"price"`
	Description  string  `json:"description" form:"description"`
	ContractorID int     `json:"contractor_id" form:"contractor_id"`
}

func ToCore(req Portfolio) portfolios.Core {
	return portfolios.Core{
		ClientName:  req.ClientName,
		Location:    req.Location,
		FinishDate:  req.FinishDate,
		Longitude:   req.Longitude,
		Latitude:    req.Latitude,
		Price:       req.Price,
		Description: req.Description,
		Contractor: portfolios.Contractor{
			ID: req.ContractorID,
		},
	}
}
