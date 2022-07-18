package business

import (
	"capstoneproject/features/portfolios"
	"fmt"
)

type portfolioUsecase struct {
	portfolioData portfolios.Data
}

func NewPortfolioBusiness(portData portfolios.Data) portfolios.Business {
	return &portfolioUsecase{
		portfolioData: portData,
	}
}

func (uc *portfolioUsecase) PostPortfolio(input portfolios.Core) (idPortfolio int, row int, err error) {
	if input.ClientName == "" || input.Latitude == 0 || input.Longitude == 0 || input.Price == 0 || input.Description == "" {
		return 0, -1, fmt.Errorf("all input must be filled")
	}
	idPortfolio, row, err = uc.portfolioData.InsertPortfolio(input)
	return idPortfolio, row, err
}