package data

import (
	"capstoneproject/features/portfolios"

	"gorm.io/gorm"
)

type mysqlPortfolioRepository struct {
	db *gorm.DB
}

func NewPortfolioRepository(conn *gorm.DB) portfolios.Data {
	return &mysqlPortfolioRepository{
		db: conn,
	}
}

func (repo *mysqlPortfolioRepository) InsertPortfolio(data portfolios.Core) (int, int, error) {
	var dataPortfolio = fromCore(data)
	result := repo.db.Create(&dataPortfolio)
	if result.Error != nil {
		return 0, 0, result.Error
	}
	return int(dataPortfolio.ID), 1, nil
}
