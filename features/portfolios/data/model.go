package data

import (
	"capstoneproject/features/portfolios"

	"gorm.io/gorm"
)

type Portfolio struct {
	gorm.Model
	ClientName     string  `json:"client_name"`
	Location       string  `json:"location"`
	FinishDate     string  `json:"finish_date"`
	Longitude      float64 `json:"longitude"`
	Latitude       float64 `json:"latitude"`
	Price          int     `json:"price"`
	Description    string  `json:"description"`
	ContractorID   uint    `json:"contractor_id"`
	Contractor     Contractor
	PortfolioImage []PortfolioImage
}

type Contractor struct {
	gorm.Model
	Portfolio []Portfolio
}

type PortfolioImage struct {
	gorm.Model
	PortfolioID uint   `json:"portfolio_id"`
	ImageURL    string `json:"image_url"`
	Portfolio   Portfolio
}

func (data *Portfolio) toCore() portfolios.Core {
	return portfolios.Core{
		ID:          int(data.ID),
		ClientName:  data.ClientName,
		Location:    data.Location,
		FinishDate:  data.FinishDate,
		Longitude:   data.Longitude,
		Latitude:    data.Latitude,
		Price:       data.Price,
		Description: data.Description,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		Contractor: portfolios.Contractor{
			ID: int(data.Contractor.ID),
		},
		PortfolioImage: toCoreListPortfolioImages(data.PortfolioImage),
	}
}

func toCoreList(data []Portfolio) []portfolios.Core {
	result := []portfolios.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core portfolios.Core) Portfolio {
	return Portfolio{
		ClientName:   core.ClientName,
		Location:     core.Location,
		FinishDate:   core.FinishDate,
		Longitude:    core.Longitude,
		Latitude:     core.Latitude,
		Price:        core.Price,
		Description:  core.Description,
		ContractorID: uint(core.Contractor.ID),
	}
}

func (data *PortfolioImage) toCorePortfolioImages() portfolios.PortfolioImage {
	return portfolios.PortfolioImage{
		ID:       int(data.ID),
		ImageURL: data.ImageURL,
		Portfolio: portfolios.Core{
			ID: int(data.PortfolioID),
		},
	}
}

func toCoreListPortfolioImages(data []PortfolioImage) []portfolios.PortfolioImage {
	result := []portfolios.PortfolioImage{}
	for key := range data {
		result = append(result, data[key].toCorePortfolioImages())
	}
	return result
}
