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

func TestGetAllHouse(t *testing.T) {
	t.Run("Test Get All House Success", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseData{})
		result, totalPage, err := houseBusiness.GetAllHouse(0, 0)
		assert.Nil(t, err)
		assert.Equal(t, []houses.Core{
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
		}, result)
		assert.Equal(t, 1, totalPage)
	})
	t.Run("Test Get All House Success When Limit is Odd", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseData{})
		result, totalPage, err := houseBusiness.GetAllHouse(1, 0)
		assert.Nil(t, err)
		assert.Equal(t, []houses.Core{
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
		}, result)
		assert.Equal(t, 20, totalPage)
	})
	t.Run("Test Get All House Success When Limit is Even", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseData{})
		result, totalPage, err := houseBusiness.GetAllHouse(2, 0)
		assert.Nil(t, err)
		assert.Equal(t, []houses.Core{
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
		}, result)
		assert.Equal(t, 10, totalPage)
	})
	t.Run("Test Get All House Failed", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseDataFailed{})
		result, totalPage, err := houseBusiness.GetAllHouse(0, 0)
		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Equal(t, 0, totalPage)
	})
}

func TestPostNewHouse(t *testing.T) {
	t.Run("Test Post New House Success", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseData{})
		var data = houses.Core{
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
	t.Run("Test Post New House Failed", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseDataFailed{})
		var data = houses.Core{
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
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
		assert.Equal(t, 0, idHouse)
	})
	t.Run("Test Post New House Failed When Title is Empty", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseDataFailed{})
		var data = houses.Core{
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
		assert.NotNil(t, err)
		assert.Equal(t, -1, row)
		assert.Equal(t, 0, idHouse)
	})
}

func TestGetHouseDetail(t *testing.T) {
	t.Run("Test Get House Detail Success", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseData{})
		result, err := houseBusiness.GetHouseDetail(1)
		assert.Nil(t, err)
		assert.Equal(t, houses.Core{
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
		}, result)
	})
	t.Run("Test Get House Detail Success", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseDataFailed{})
		result, err := houseBusiness.GetHouseDetail(1)
		assert.NotNil(t, err)
		assert.Equal(t, houses.Core{}, result)
	})
}

func TestGetMyListHouse(t *testing.T) {
	t.Run("Test Get My List House Success", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseData{})
		result, totalPage, err := houseBusiness.GetMyListHouse(1, 0, 0)
		assert.Nil(t, err)
		assert.Equal(t, []houses.Core{
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
		}, result)
		assert.Equal(t, 1, totalPage)
	})
	t.Run("Test Get My List House Success When Limit is Odd", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseData{})
		result, totalPage, err := houseBusiness.GetMyListHouse(1, 1, 0)
		assert.Nil(t, err)
		assert.Equal(t, []houses.Core{
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
		}, result)
		assert.Equal(t, 5, totalPage)
	})
	t.Run("Test Get My List House Success When Limit is Even", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseData{})
		result, totalPage, err := houseBusiness.GetMyListHouse(1, 2, 0)
		assert.Nil(t, err)
		assert.Equal(t, []houses.Core{
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
		}, result)
		assert.Equal(t, 3, totalPage)
	})
	t.Run("Test Get My List House Failed", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseDataFailed{})
		result, totalPage, err := houseBusiness.GetMyListHouse(1, 0, 0)
		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Equal(t, 0, totalPage)
	})
}

func TestPutHouse(t *testing.T) {
	t.Run("Test Put House Success", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseData{})
		var data = houses.Core{
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
		row, err := houseBusiness.PutHouse(1, data)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})
	t.Run("Test Put House Failed", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseDataFailed{})
		var data = houses.Core{
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
		row, err := houseBusiness.PutHouse(1, data)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})
}

func TestDeleteHouse(t *testing.T) {
	t.Run("Test Delete House Success", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseData{})
		row, err := houseBusiness.DeleteHouse(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})
	t.Run("Test Delete House Failed", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseDataFailed{})
		row, err := houseBusiness.DeleteHouse(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})
}

func TestGetSearchHouse(t *testing.T) {
	t.Run("Test Get Search House Success", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseData{})
		result, totalPage, err := houseBusiness.GetSearchHouse("Rumah", "Bogor", "100000000", "3000000000", 0, 0)
		assert.Nil(t, err)
		assert.Equal(t, []houses.Core{
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
		}, result)
		assert.Equal(t, 1, totalPage)
	})
	t.Run("Test Get Search House Success When Limit is Odd", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseData{})
		result, totalPage, err := houseBusiness.GetSearchHouse("Rumah", "Bogor", "100000000", "3000000000", 1, 0)
		assert.Nil(t, err)
		assert.Equal(t, []houses.Core{
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
		}, result)
		assert.Equal(t, 10, totalPage)
	})
	t.Run("Test Get Search House Success When Limit is Even", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseData{})
		result, totalPage, err := houseBusiness.GetSearchHouse("Rumah", "Bogor", "100000000", "3000000000", 6, 0)
		assert.Nil(t, err)
		assert.Equal(t, []houses.Core{
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
		}, result)
		assert.Equal(t, 2, totalPage)
	})
	t.Run("Test Get Search House Success When Location Not Empty", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseData{})
		result, totalPage, err := houseBusiness.GetSearchHouse("Rumah", "Bogor", "", "", 0, 0)
		assert.Nil(t, err)
		assert.Equal(t, []houses.Core{
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
		}, result)
		assert.Equal(t, 1, totalPage)
	})
	t.Run("Test Get Search House Success When Min Price Not Empty", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseData{})
		result, totalPage, err := houseBusiness.GetSearchHouse("Rumah", "", "100000000", "", 0, 0)
		assert.Nil(t, err)
		assert.Equal(t, []houses.Core{
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
		}, result)
		assert.Equal(t, 1, totalPage)
	})
	t.Run("Test Get Search House Success When Max Price Not Empty", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseData{})
		result, totalPage, err := houseBusiness.GetSearchHouse("Rumah", "", "", "300000000", 0, 0)
		assert.Nil(t, err)
		assert.Equal(t, []houses.Core{
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
		}, result)
		assert.Equal(t, 1, totalPage)
	})
	t.Run("Test Get Search House Success When Min and Max Price Not Empty", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseData{})
		result, totalPage, err := houseBusiness.GetSearchHouse("Rumah", "", "100000000", "300000000", 0, 0)
		assert.Nil(t, err)
		assert.Equal(t, []houses.Core{
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
		}, result)
		assert.Equal(t, 1, totalPage)
	})
	t.Run("Test Get Search House Failed", func(t *testing.T) {
		houseBusiness := NewHouseBusiness(mockHouseDataFailed{})
		result, totalPage, err := houseBusiness.GetSearchHouse("Rumah", "", "", "", 0, 0)
		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Equal(t, 0, totalPage)
	})
}
