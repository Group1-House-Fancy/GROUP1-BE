package request

import portfolioimages "capstoneproject/features/portfolioImages"

type PortfolioImage struct {
	ImageURL    string `json:"image_url" form:"image_url"`
	PortfolioID uint   `json:"house_id" form:"house_id"`
}

func ToCore(req PortfolioImage) portfolioimages.Core {
	return portfolioimages.Core{
		ImageURL: req.ImageURL,
		Portfolio: portfolioimages.Portfolio{
			ID: int(req.PortfolioID),
		},
	}
}
