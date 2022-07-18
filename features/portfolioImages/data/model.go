package data

import (
	portfolioimages "capstoneproject/features/portfolioImages"

	"gorm.io/gorm"
)

type PortfolioImage struct {
	gorm.Model
	ImageURL    string `json:"image_url"`
	PortfolioID uint   `json:"portfolio_id"`
}

func (data *PortfolioImage) toCore() portfolioimages.Core {
	return portfolioimages.Core{
		ID:        int(data.ID),
		ImageURL:  data.ImageURL,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		Portfolio: portfolioimages.Portfolio{
			ID: int(data.PortfolioID),
		},
	}
}

func toCoreList(data []PortfolioImage) []portfolioimages.Core {
	result := []portfolioimages.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core portfolioimages.Core) PortfolioImage {
	return PortfolioImage{
		ImageURL:    core.ImageURL,
		PortfolioID: uint(core.Portfolio.ID),
	}
}
