package business

import (
	"capstoneproject/features/houses"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockHouseData struct{}

func (mock mockHouseData) SelectAllHouse(limit, offset int) ([]houses.Core, error) {
	return []houses.Core{
		{
			ID:           1,
			Title:        "Rumah Taman Darmo Permai Utara",
			Price:        250000000,
			Location:     "Sentul, Bogor",
			Longitude:    106.82181,
			Latitude:     -6.193125,
			SurfaceArea:  64,
			BuildingArea: 124,
			Bathroom:     2,
			Bedroom:      4,
			Certificate:  "SHM",
			Status:       "Available",
			Description:  "Rumah Dijual di Bogor RUMAH TAMAN DARMO PERMAI UTARA  LT 135 LB 90 KT 3 KM 2 2LANTAI AC 2 UNIT 2200W SUDAH RENOV HARGA 50JT/TH (NETT)",
			HouseImage: []houses.HouseImage{
				{ID: 1, ImageURL: "https://storage.googleapis.com/bucket-project-capstone/tampak-depan.jpg"},
				{ID: 2, ImageURL: "https://storage.googleapis.com/bucket-project-capstone/tampak-samping.jpg"},
			},
			User: houses.User{
				ID:          1,
				FullName:    "Adi Setiawan",
				Email:       "adi@mail.com",
				PhoneNumber: "085345654123",
				Address:     "Yogyakarta",
				ImageURL:    "https://storage.googleapis.com/bucket-project-capstone/default_profile.png",
			},
		},
	}, nil
}

func (mock mockHouseData) InsertNewHouse(data houses.Core) (int, int, error) {
	return 1, 1, nil
}

func (mock mockHouseData) SelectHouseByIdHouse(idHouse int) (houses.Core, error) {
	return houses.Core{
		ID:           1,
		Title:        "Rumah Taman Darmo Permai Utara",
		Price:        250000000,
		Location:     "Sentul, Bogor",
		Longitude:    106.82181,
		Latitude:     -6.193125,
		SurfaceArea:  64,
		BuildingArea: 124,
		Bathroom:     2,
		Bedroom:      4,
		Certificate:  "SHM",
		Status:       "Available",
		Description:  "Rumah Dijual di Bogor RUMAH TAMAN DARMO PERMAI UTARA  LT 135 LB 90 KT 3 KM 2 2LANTAI AC 2 UNIT 2200W SUDAH RENOV HARGA 50JT/TH (NETT)",
		HouseImage: []houses.HouseImage{
			{ID: 1, ImageURL: "https://storage.googleapis.com/bucket-project-capstone/tampak-depan.jpg"},
			{ID: 2, ImageURL: "https://storage.googleapis.com/bucket-project-capstone/tampak-samping.jpg"},
		},
		User: houses.User{
			ID:          1,
			FullName:    "Adi Setiawan",
			Email:       "adi@mail.com",
			PhoneNumber: "085345654123",
			Address:     "Yogyakarta",
			ImageURL:    "https://storage.googleapis.com/bucket-project-capstone/default_profile.png",
		},
	}, nil
}

func (mock mockHouseData) SelectHouseByIdUser(idUser, limit, offset int) ([]houses.Core, error) {
	return []houses.Core{
		{
			ID:           1,
			Title:        "Rumah Taman Darmo Permai Utara",
			Price:        250000000,
			Location:     "Sentul, Bogor",
			Longitude:    106.82181,
			Latitude:     -6.193125,
			SurfaceArea:  64,
			BuildingArea: 124,
			Bathroom:     2,
			Bedroom:      4,
			Certificate:  "SHM",
			Status:       "Available",
			Description:  "Rumah Dijual di Bogor RUMAH TAMAN DARMO PERMAI UTARA  LT 135 LB 90 KT 3 KM 2 2LANTAI AC 2 UNIT 2200W SUDAH RENOV HARGA 50JT/TH (NETT)",
			HouseImage: []houses.HouseImage{
				{ID: 1, ImageURL: "https://storage.googleapis.com/bucket-project-capstone/tampak-depan.jpg"},
				{ID: 2, ImageURL: "https://storage.googleapis.com/bucket-project-capstone/tampak-samping.jpg"},
			},
			User: houses.User{
				ID:          1,
				FullName:    "Adi Setiawan",
				Email:       "adi@mail.com",
				PhoneNumber: "085345654123",
				Address:     "Yogyakarta",
				ImageURL:    "https://storage.googleapis.com/bucket-project-capstone/default_profile.png",
			},
		}, {
			ID:           2,
			Title:        "Rumah Taman Darmo Permai Utara",
			Price:        250000000,
			Location:     "Sentul, Bogor",
			Longitude:    106.82183,
			Latitude:     -6.193122,
			SurfaceArea:  64,
			BuildingArea: 124,
			Bathroom:     2,
			Bedroom:      4,
			Certificate:  "SHM",
			Status:       "Available",
			Description:  "Rumah Dijual di Bogor RUMAH TAMAN DARMO PERMAI UTARA  LT 135 LB 90 KT 3 KM 2 2LANTAI AC 2 UNIT 2200W SUDAH RENOV HARGA 50JT/TH (NETT)",
			HouseImage: []houses.HouseImage{
				{ID: 1, ImageURL: "https://storage.googleapis.com/bucket-project-capstone/tampak-depan.jpg"},
				{ID: 2, ImageURL: "https://storage.googleapis.com/bucket-project-capstone/tampak-samping.jpg"},
			},
			User: houses.User{
				ID:          1,
				FullName:    "Adi Setiawan",
				Email:       "adi@mail.com",
				PhoneNumber: "085345654123",
				Address:     "Yogyakarta",
				ImageURL:    "https://storage.googleapis.com/bucket-project-capstone/default_profile.png",
			},
		},
	}, nil
}

func (mock mockHouseData) UpdateHouse(idHouse int, data houses.Core) (int, error) {
	return 1, nil
}

func (mock mockHouseData) DeleteHouse(idHouse int) (int, error) {
	return 1, nil
}

func (mock mockHouseData) SelectSearchHouse(keywords string, limit, offset int) ([]houses.Core, error) {
	return []houses.Core{
		{
			ID:           1,
			Title:        "Rumah Taman Darmo Permai Utara",
			Price:        250000000,
			Location:     "Sentul, Bogor",
			Longitude:    106.82181,
			Latitude:     -6.193125,
			SurfaceArea:  64,
			BuildingArea: 124,
			Bathroom:     2,
			Bedroom:      4,
			Certificate:  "SHM",
			Status:       "Available",
			Description:  "Rumah Dijual di Bogor RUMAH TAMAN DARMO PERMAI UTARA  LT 135 LB 90 KT 3 KM 2 2LANTAI AC 2 UNIT 2200W SUDAH RENOV HARGA 50JT/TH (NETT)",
			HouseImage: []houses.HouseImage{
				{ID: 1, ImageURL: "https://storage.googleapis.com/bucket-project-capstone/tampak-depan.jpg"},
				{ID: 2, ImageURL: "https://storage.googleapis.com/bucket-project-capstone/tampak-samping.jpg"},
			},
			User: houses.User{
				ID:          1,
				FullName:    "Adi Setiawan",
				Email:       "adi@mail.com",
				PhoneNumber: "085345654123",
				Address:     "Yogyakarta",
				ImageURL:    "https://storage.googleapis.com/bucket-project-capstone/default_profile.png",
			},
		},
	}, nil
}

func (mock mockHouseData) CountHouseData() (int, error) {
	return 20, nil
}

func (mock mockHouseData) CountMyListHouseData(idUser int) (int, error) {
	return 5, nil
}

func (mock mockHouseData) CountSearchHouseData(query string) (int, error) {
	return 10, nil
}

type mockHouseDataFailed struct{}

func (mock mockHouseDataFailed) SelectAllHouse(limit, offset int) ([]houses.Core, error) {
	return nil, fmt.Errorf("error to get all data")
}

func (mock mockHouseDataFailed) InsertNewHouse(data houses.Core) (int, int, error) {
	return 0, 0, fmt.Errorf("error to insert data")
}

func (mock mockHouseDataFailed) SelectHouseByIdHouse(idHouse int) (houses.Core, error) {
	return houses.Core{}, fmt.Errorf("error to get data")
}

func (mock mockHouseDataFailed) SelectHouseByIdUser(idUser, limit, offset int) ([]houses.Core, error) {
	return nil, fmt.Errorf("error to get all data")
}

func (mock mockHouseDataFailed) UpdateHouse(idHouse int, data houses.Core) (int, error) {
	return 0, fmt.Errorf("error to update data")
}

func (mock mockHouseDataFailed) DeleteHouse(idHouse int) (int, error) {
	return 0, fmt.Errorf("error to delete data")
}

func (mock mockHouseDataFailed) SelectSearchHouse(keywords string, limit, offset int) ([]houses.Core, error) {
	return nil, fmt.Errorf("error to get all data")
}

func (mock mockHouseDataFailed) CountHouseData() (int, error) {
	return 0, fmt.Errorf("error to count data")
}

func (mock mockHouseDataFailed) CountMyListHouseData(idUser int) (int, error) {
	return 0, fmt.Errorf("error to count data")
}

func (mock mockHouseDataFailed) CountSearchHouseData(query string) (int, error) {
	return 0, fmt.Errorf("error to count data")
}

func TestPostNewHouse(t *testing.T) {
	t.Run("Test Post New House Image Success", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseData{})
		var data = houses.Core{
			ID:           1,
			Title:        "Rumah Taman Darmo Permai Utara",
			Price:        250000000,
			Location:     "Sentul, Bogor",
			Longitude:    106.82181,
			Latitude:     -6.193125,
			SurfaceArea:  64,
			BuildingArea: 124,
			Bathroom:     2,
			Bedroom:      4,
			Certificate:  "SHM",
			Status:       "Available",
			Description:  "Rumah Dijual di Bogor RUMAH TAMAN DARMO PERMAI UTARA  LT 135 LB 90 KT 3 KM 2 2LANTAI AC 2 UNIT 2200W SUDAH RENOV HARGA 50JT/TH (NETT)",
			User: houses.User{
				ID: 1,
			},
		}
		idHouse, row, err := houseBusiness.PostNewHouse(data)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
		assert.Equal(t, 1, idHouse)
	})
}
