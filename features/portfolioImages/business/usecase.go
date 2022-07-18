package business

import (
	portfolioimages "capstoneproject/features/portfolioImages"
	"fmt"
)

type portfolioImageUsecase struct {
	portfolioImageData portfolioimages.Data
}

func NewPortfolioImageBusiness(portfolioImgData portfolioimages.Data) portfolioimages.Business {
	return &portfolioImageUsecase{
		portfolioImageData: portfolioImgData,
	}
}

func (uc *portfolioImageUsecase) PostNewPortfolioImage(data portfolioimages.Core) (row int, err error) {
	if data.ImageURL == "" || data.Portfolio.ID == 0 {
		return -1, fmt.Errorf("all input must be filled")
	}
	row, err = uc.portfolioImageData.InsertNewImage(data)
	return row, err
}

func (uc *portfolioImageUsecase) DeleteImage(idImage int) (row int, err error) {
	row, err = uc.portfolioImageData.DeleteImage(idImage)
	return row, err
}
