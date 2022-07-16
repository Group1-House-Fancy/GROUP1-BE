package business

import (
	"capstoneproject/features/houses"
	"fmt"
)

type houseUsecase struct {
	houseData houses.Data
}

func NewHouseBusiness(housData houses.Data) houses.Business {
	return &houseUsecase{
		houseData: housData,
	}
}

func (uc *houseUsecase) GetAllHouse(limit, offset int) (resp []houses.Core, totalPage int, err error) {
	resp, err = uc.houseData.SelectAllHouse(limit, offset)
	total := len(resp)
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
	return resp, totalPage, err
}

func (uc *houseUsecase) PostNewHouse(input houses.Core) (idHouse int, row int, err error) {
	if input.Title == "" || input.Latitude == 0 || input.Longitude == 0 {
		return 0, -1, fmt.Errorf("all input must be filled")
	}
	idHouse, row, err = uc.houseData.InsertNewHouse(input)
	return idHouse, row, err
}

func (uc *houseUsecase) GetHouseDetail(idHouse int) (resp houses.Core, err error) {
	resp, err = uc.houseData.SelectHouseByIdHouse(idHouse)
	return resp, err
}

func (uc *houseUsecase) GetMyListHouse(idUser, limit, offset int) (resp []houses.Core, totalPage int, err error) {
	resp, err = uc.houseData.SelectHouseByIdUser(idUser, limit, offset)
	total := len(resp)
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
	return resp, totalPage, err
}
func (uc *houseUsecase) PutHouse(idHouse int, input houses.Core) (row int, err error) {
	if input.Title == "" || input.Latitude == 0 || input.Longitude == 0 {
		return -1, fmt.Errorf("all input must be filled")
	}
	row, err = uc.houseData.UpdateHouse(idHouse, input)
	return row, err
}
