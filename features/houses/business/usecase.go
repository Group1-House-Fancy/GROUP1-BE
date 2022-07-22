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
	total, _ := uc.houseData.CountHouseData()
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
	if input.Title == "" || input.Latitude == 0 || input.Longitude == 0 || input.Location == "" || input.BuildingArea == 0 || input.SurfaceArea == 0 || input.Price == 0 || input.Bedroom == 0 || input.Bathroom == 0 || input.Certificate == "" {
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
	total, _ := uc.houseData.CountMyListHouseData(idUser)
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
	if input.Title == "" || input.Latitude == 0 || input.Longitude == 0 || input.Location == "" || input.BuildingArea == 0 || input.SurfaceArea == 0 || input.Price == 0 || input.Bedroom == 0 || input.Bathroom == 0 || input.Certificate == "" {
		return -1, fmt.Errorf("all input must be filled")
	}
	row, err = uc.houseData.UpdateHouse(idHouse, input)
	return row, err
}

func (uc *houseUsecase) DeleteHouse(idHouse int) (row int, err error) {
	row, err = uc.houseData.DeleteHouse(idHouse)
	return row, err
}

func (uc *houseUsecase) GetSearchHouse(keywords, location, minPrice, maxPrice string, limit, offset int) (resp []houses.Core, totalPage int, err error) {
	var query string = ""
	if minPrice != "" || maxPrice != "" {
		query += "( "
		if minPrice != "" {
			query += "price >= " + minPrice
		} else if maxPrice != "" {
			query += "price <= " + maxPrice
		} else if minPrice != "" && maxPrice != "" {
			query += "price BETWEEN " + minPrice + " AND " + maxPrice + ""
		}
		query += " )"
	}
	if location != "" {
		if minPrice != "" || maxPrice != "" {
			query += " AND ( "
		} else {
			query += "( "
		}
		query += "location LIKE '%" + location + "%'"
		query += " )"
	}
	if minPrice != "" || maxPrice != "" || location != "" {
		query += " AND (title LIKE '%" + keywords + "%' OR location LIKE '%" + keywords + "%' OR description LIKE '%" + keywords + "%')"
	} else {
		query += "(title LIKE '%" + keywords + "%' OR location LIKE '%" + keywords + "%' OR description LIKE '%" + keywords + "%')"
	}
	resp, err = uc.houseData.SelectSearchHouse(query, limit, offset)
	total, _ := uc.houseData.CountSearchHouseData(query)
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
