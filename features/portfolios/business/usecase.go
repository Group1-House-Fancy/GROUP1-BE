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

func (uc *portfolioUsecase) GetAllPortfolio(idCtr, limit, offset int) (data []portfolios.Core, totalPage int, err error) {
	data, err = uc.portfolioData.SelectAllPortfolio(idCtr, limit, offset)
	total := len(data)
	if total == 0 {
		totalPage = 0
	} else {
		if limit == 0 {
			limit = total
		}
		if total%limit != 0 {
			totalPage = (total / limit) + 1
		} else {
			totalPage = total / limit
		}
	}
	return data, totalPage, err
}

func (uc *portfolioUsecase) GetPortfolio(idPrtf int) (data portfolios.Core, err error) {
	data, err = uc.portfolioData.SelectPortfolio(idPrtf)
	return data, err
}
