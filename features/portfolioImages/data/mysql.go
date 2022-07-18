package data

import (
	portfolioimages "capstoneproject/features/portfolioImages"

	"gorm.io/gorm"
)

type mysqlPortfolioImageRepository struct {
	db *gorm.DB
}

func NewPortfolioImageRepository(conn *gorm.DB) portfolioimages.Data {
	return &mysqlPortfolioImageRepository{
		db: conn,
	}
}

func (repo *mysqlPortfolioImageRepository) InsertNewImage(data portfolioimages.Core) (int, error) {
	var dataImage = fromCore(data)
	result := repo.db.Create(&dataImage)
	if result.Error != nil {
		return 0, result.Error
	}
	return 1, nil
}

func (repo *mysqlPortfolioImageRepository) DeleteImage(idImage int) (int, error) {
	result := repo.db.Where("id = ?", idImage).Delete(&PortfolioImage{})
	if result.Error != nil {
		return 0, result.Error
	}
	return 1, nil
}
