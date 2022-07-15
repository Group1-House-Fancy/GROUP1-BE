package business

import "capstoneproject/features/houses"

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
