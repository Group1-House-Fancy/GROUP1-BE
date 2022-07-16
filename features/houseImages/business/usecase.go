package business

import (
	houseimages "capstoneproject/features/houseImages"
	"fmt"
)

type houseImageUsecase struct {
	houseImageData houseimages.Data
}

func NewHouseImageBusiness(houseImgData houseimages.Data) houseimages.Business {
	return &houseImageUsecase{
		houseImageData: houseImgData,
	}
}

func (uc *houseImageUsecase) PostNewHouseImage(data houseimages.Core) (row int, err error) {
	if data.ImageURL == "" || data.House.ID == 0 {
		return -1, fmt.Errorf("all input must be filled")
	}
	row, err = uc.houseImageData.InsertNewImages(data)
	return row, err
}
