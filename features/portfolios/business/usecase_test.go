package business

import (
	"capstoneproject/features/portfolios"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockPortfolioData struct{}

func (mock mockPortfolioData) InsertPortfolio(data portfolios.Core) (int, int, error) {
	return 1, 1, nil
}

func (mock mockPortfolioData) SelectAllPortfolio(idCtr, limit, offset int) ([]portfolios.Core, error) {
	return []portfolios.Core{
		{
			ID:          1,
			ClientName:  "Azka",
			Location:    "Pekalongan",
			FinishDate:  "2 Mei 2021",
			Longitude:   106.82181,
			Latitude:    -6.193125,
			Price:       20000000,
			Description: "renovasi rumah",
			Contractor: portfolios.Contractor{
				ID: 1,
			},
			PortfolioImage: []portfolios.PortfolioImage{
				{ID: 1, ImageURL: "https://storage.googleapis.com/bucket-project-capstone/ad0a0d9e-8bfa-4e72-85ef-ab6e79336a88.jpg"},
				{ID: 2, ImageURL: "https://storage.googleapis.com/bucket-project-capstone/ad0a0d9e-8bfa-4e72-85ef-ab6e79336a88.jpg"},
			},
		},
	}, nil
}

func (mock mockPortfolioData) SelectPortfolio(idPrtf int) (portfolios.Core, error) {
	return portfolios.Core{
		ID:          2,
		ClientName:  "Anhar Kusuma",
		Location:    "Jawa Barat",
		FinishDate:  "08 Desember 2021",
		Longitude:   106.82181,
		Latitude:    -6.193125,
		Price:       100000000,
		Description: "Renovasi Rumah",
		PortfolioImage: []portfolios.PortfolioImage{
			{ID: 1, ImageURL: "https://storage.googleapis.com/bucket-project-3/portfolio1.png"},
			{ID: 2, ImageURL: "https://storage.googleapis.com/bucket-project-3/portfolio2.png"},
			{ID: 3, ImageURL: "https://storage.googleapis.com/bucket-project-3/portfolio3.png"},
			{ID: 4, ImageURL: "https://storage.googleapis.com/bucket-project-3/portfolio4.png"},
			{ID: 5, ImageURL: "https://storage.googleapis.com/bucket-project-3/portfolio5.png"},
		},
	}, nil
}

func (mock mockPortfolioData) UpdatePortfolio(idPrtf int, data portfolios.Core) (int, int, error) {
	return 1, 1, nil
}

func (mock mockPortfolioData) DeletePortfolio(idPrtf int) (int, error) {
	return 1, nil
}

func (mock mockPortfolioData) CountPortfolioData(idCtr int) (int, error) {
	return 6, nil
}

type mockPortfolioDataFailed struct{}

func (mock mockPortfolioDataFailed) InsertPortfolio(data portfolios.Core) (int, int, error) {
	return 0, 0, fmt.Errorf("failed to insert data")
}

func (mock mockPortfolioDataFailed) SelectAllPortfolio(idCtr, limit, offset int) ([]portfolios.Core, error) {
	return nil, fmt.Errorf("failed to get all data")
}

func (mock mockPortfolioDataFailed) SelectPortfolio(idPrtf int) (portfolios.Core, error) {
	return portfolios.Core{}, fmt.Errorf("failed to get data")
}

func (mock mockPortfolioDataFailed) UpdatePortfolio(idPrtf int, data portfolios.Core) (int, int, error) {
	return 0, 0, fmt.Errorf("failed to update data")
}

func (mock mockPortfolioDataFailed) DeletePortfolio(idPrtf int) (int, error) {
	return 0, fmt.Errorf("failed to delete data")
}

func (mock mockPortfolioDataFailed) CountPortfolioData(idCtr int) (int, error) {
	return 0, fmt.Errorf("failed to count data")
}

func TestPostPortfolio(t *testing.T) {
	t.Run("Test Post Portfolio Success", func(t *testing.T) {
		portfolioBusiness := NewPortfolioBusiness(mockPortfolioData{})
		var data = portfolios.Core{
			ClientName:  "Aldi",
			Location:    "Pekalongan",
			FinishDate:  "2 Mei 2021",
			Longitude:   106.821810,
			Latitude:    -6.193125,
			Price:       20000000,
			Description: "renovasi rumah",
			Contractor: portfolios.Contractor{
				ID: 3,
			},
		}
		idPortfolio, row, err := portfolioBusiness.PostPortfolio(data)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
		assert.Equal(t, 1, idPortfolio)
	})
	t.Run("Test Post Portfolio Failed", func(t *testing.T) {
		portfolioBusiness := NewPortfolioBusiness(mockPortfolioDataFailed{})
		var data = portfolios.Core{
			ClientName:  "Aldi",
			Location:    "Pekalongan",
			FinishDate:  "2 Mei 2021",
			Longitude:   106.821810,
			Latitude:    -6.193125,
			Price:       20000000,
			Description: "renovasi rumah",
			Contractor: portfolios.Contractor{
				ID: 3,
			},
		}
		idPortfolio, row, err := portfolioBusiness.PostPortfolio(data)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
		assert.Equal(t, 0, idPortfolio)
	})
	t.Run("Test Post Portfolio Failed When ClientName is Empty", func(t *testing.T) {
		portfolioBusiness := NewPortfolioBusiness(mockPortfolioDataFailed{})
		var data = portfolios.Core{
			Location:    "Pekalongan",
			FinishDate:  "2 Mei 2021",
			Longitude:   106.821810,
			Latitude:    -6.193125,
			Price:       20000000,
			Description: "renovasi rumah",
			Contractor: portfolios.Contractor{
				ID: 3,
			},
		}
		idPortfolio, row, err := portfolioBusiness.PostPortfolio(data)
		assert.NotNil(t, err)
		assert.Equal(t, -1, row)
		assert.Equal(t, 0, idPortfolio)
	})
}

func TestGetAllPortfolio(t *testing.T) {
	t.Run("Test Get All Portfolio Success", func(t *testing.T) {
		portfolioBusiness := NewPortfolioBusiness(mockPortfolioData{})
		data, totalPage, err := portfolioBusiness.GetAllPortfolio(1, 0, 0)
		assert.Nil(t, err)
		assert.Equal(t, []portfolios.Core{
			{
				ID:          1,
				ClientName:  "Azka",
				Location:    "Pekalongan",
				FinishDate:  "2 Mei 2021",
				Longitude:   106.82181,
				Latitude:    -6.193125,
				Price:       20000000,
				Description: "renovasi rumah",
				Contractor: portfolios.Contractor{
					ID: 1,
				},
				PortfolioImage: []portfolios.PortfolioImage{
					{ID: 1, ImageURL: "https://storage.googleapis.com/bucket-project-capstone/ad0a0d9e-8bfa-4e72-85ef-ab6e79336a88.jpg"},
					{ID: 2, ImageURL: "https://storage.googleapis.com/bucket-project-capstone/ad0a0d9e-8bfa-4e72-85ef-ab6e79336a88.jpg"},
				},
			},
		}, data)
		assert.Equal(t, 1, totalPage)
	})
	t.Run("Test Get All Portfolio Success When Limit is Odd", func(t *testing.T) {
		portfolioBusiness := NewPortfolioBusiness(mockPortfolioData{})
		data, totalPage, err := portfolioBusiness.GetAllPortfolio(1, 1, 0)
		assert.Nil(t, err)
		assert.Equal(t, []portfolios.Core{
			{
				ID:          1,
				ClientName:  "Azka",
				Location:    "Pekalongan",
				FinishDate:  "2 Mei 2021",
				Longitude:   106.82181,
				Latitude:    -6.193125,
				Price:       20000000,
				Description: "renovasi rumah",
				Contractor: portfolios.Contractor{
					ID: 1,
				},
				PortfolioImage: []portfolios.PortfolioImage{
					{ID: 1, ImageURL: "https://storage.googleapis.com/bucket-project-capstone/ad0a0d9e-8bfa-4e72-85ef-ab6e79336a88.jpg"},
					{ID: 2, ImageURL: "https://storage.googleapis.com/bucket-project-capstone/ad0a0d9e-8bfa-4e72-85ef-ab6e79336a88.jpg"},
				},
			},
		}, data)
		assert.Equal(t, 6, totalPage)
	})
	t.Run("Test Get All Portfolio Success When Limit is Even", func(t *testing.T) {
		portfolioBusiness := NewPortfolioBusiness(mockPortfolioData{})
		data, totalPage, err := portfolioBusiness.GetAllPortfolio(1, 2, 0)
		assert.Nil(t, err)
		assert.Equal(t, []portfolios.Core{
			{
				ID:          1,
				ClientName:  "Azka",
				Location:    "Pekalongan",
				FinishDate:  "2 Mei 2021",
				Longitude:   106.82181,
				Latitude:    -6.193125,
				Price:       20000000,
				Description: "renovasi rumah",
				Contractor: portfolios.Contractor{
					ID: 1,
				},
				PortfolioImage: []portfolios.PortfolioImage{
					{ID: 1, ImageURL: "https://storage.googleapis.com/bucket-project-capstone/ad0a0d9e-8bfa-4e72-85ef-ab6e79336a88.jpg"},
					{ID: 2, ImageURL: "https://storage.googleapis.com/bucket-project-capstone/ad0a0d9e-8bfa-4e72-85ef-ab6e79336a88.jpg"},
				},
			},
		}, data)
		assert.Equal(t, 3, totalPage)
	})
	t.Run("Test Get All Portfolio Failed", func(t *testing.T) {
		portfolioBusiness := NewPortfolioBusiness(mockPortfolioDataFailed{})
		data, totalPage, err := portfolioBusiness.GetAllPortfolio(1, 0, 0)
		assert.NotNil(t, err)
		assert.Nil(t, data)
		assert.Equal(t, 0, totalPage)
	})
	t.Run("Test Get All Portfolio Failed When Contractor ID is empty", func(t *testing.T) {
		portfolioBusiness := NewPortfolioBusiness(mockPortfolioDataFailed{})
		data, totalPage, err := portfolioBusiness.GetAllPortfolio(0, 0, 0)
		assert.NotNil(t, err)
		assert.Nil(t, data)
		assert.NotNil(t, 0, totalPage)
	})
}

func TestGetPortfolio(t *testing.T) {
	t.Run("Test Get Portfolio Success", func(t *testing.T) {
		portfolioBusiness := NewPortfolioBusiness(mockPortfolioData{})
		result, err := portfolioBusiness.GetPortfolio(1)
		assert.Nil(t, err)
		assert.Equal(t, portfolios.Core{
			ID:          2,
			ClientName:  "Anhar Kusuma",
			Location:    "Jawa Barat",
			FinishDate:  "08 Desember 2021",
			Longitude:   106.82181,
			Latitude:    -6.193125,
			Price:       100000000,
			Description: "Renovasi Rumah",
			PortfolioImage: []portfolios.PortfolioImage{
				{ID: 1, ImageURL: "https://storage.googleapis.com/bucket-project-3/portfolio1.png"},
				{ID: 2, ImageURL: "https://storage.googleapis.com/bucket-project-3/portfolio2.png"},
				{ID: 3, ImageURL: "https://storage.googleapis.com/bucket-project-3/portfolio3.png"},
				{ID: 4, ImageURL: "https://storage.googleapis.com/bucket-project-3/portfolio4.png"},
				{ID: 5, ImageURL: "https://storage.googleapis.com/bucket-project-3/portfolio5.png"},
			},
		}, result)
	})
	t.Run("Test Get Portfolio Failed", func(t *testing.T) {
		portfolioBusiness := NewPortfolioBusiness(mockPortfolioDataFailed{})
		result, err := portfolioBusiness.GetPortfolio(1)
		assert.NotNil(t, err)
		assert.Equal(t, portfolios.Core{}, result)
	})
}

func TestPutPortfolio(t *testing.T) {
	t.Run("Test Put Portfolio Success", func(t *testing.T) {
		portfolioBusiness := NewPortfolioBusiness(mockPortfolioData{})
		var data = portfolios.Core{
			ClientName:  "Aldi",
			Location:    "Pekalongan",
			FinishDate:  "2 Mei 2021",
			Longitude:   106.821810,
			Latitude:    -6.193125,
			Price:       20000000,
			Description: "renovasi rumah",
			Contractor: portfolios.Contractor{
				ID: 3,
			},
		}
		row, resData, err := portfolioBusiness.PutPortfolio(1, data)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
		assert.Equal(t, 1, resData)
	})
	t.Run("Test Put Portfolio Failed", func(t *testing.T) {
		portfolioBusiness := NewPortfolioBusiness(mockPortfolioDataFailed{})
		var data = portfolios.Core{
			ClientName:  "Aldi",
			Location:    "Pekalongan",
			FinishDate:  "2 Mei 2021",
			Longitude:   106.821810,
			Latitude:    -6.193125,
			Price:       20000000,
			Description: "renovasi rumah",
			Contractor: portfolios.Contractor{
				ID: 3,
			},
		}
		row, resData, err := portfolioBusiness.PutPortfolio(1, data)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
		assert.Equal(t, 0, resData)
	})
	t.Run("Test Put Portfolio Failed When Client Name Empty", func(t *testing.T) {
		portfolioBusiness := NewPortfolioBusiness(mockPortfolioDataFailed{})
		var data = portfolios.Core{
			Location:    "Pekalongan",
			FinishDate:  "2 Mei 2021",
			Longitude:   106.821810,
			Latitude:    -6.193125,
			Price:       20000000,
			Description: "renovasi rumah",
			Contractor: portfolios.Contractor{
				ID: 3,
			},
		}
		row, resData, err := portfolioBusiness.PutPortfolio(0, data)
		assert.NotNil(t, err)
		assert.Equal(t, -1, row)
		assert.Equal(t, 0, resData)
	})
}

func TestDeletePortfolio(t *testing.T) {
	t.Run("Test Delete Portfolio Success", func(t *testing.T) {
		portfolioBusiness := NewPortfolioBusiness(mockPortfolioData{})
		row, err := portfolioBusiness.DeletePortfolio(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})
	t.Run("Test Delete Portfolio Failed", func(t *testing.T) {
		portfolioBusiness := NewPortfolioBusiness(mockPortfolioDataFailed{})
		row, err := portfolioBusiness.DeletePortfolio(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})
}
