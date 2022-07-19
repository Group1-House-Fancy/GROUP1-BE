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

func (repo *mysqlPortfolioRepository) SelectAllPortfolio(idCtr, limit, offset int) ([]portfolios.Core, error) {
	var dataPrt []Portfolio
	result := repo.db.Preload("Contractor").Preload("PortfolioImage").Where("contractor_id = ?", idCtr).Limit(limit).Offset(offset).Find(&dataPrt)
	if result.Error != nil {
		return nil, result.Error
	}
	return toCoreList(dataPrt), nil
}

func (repo *mysqlPortfolioRepository) SelectPortfolio(idPrtf int) (portfolios.Core, error) {
	var dataPrtf Portfolio
	result := repo.db.Preload("Contractor").Preload("PortfolioImage").Where("id = ?", idPrtf).First(&dataPrtf)
	if result.Error != nil {
		return portfolios.Core{}, result.Error
	}
	return dataPrtf.toCore(), nil
}

func (repo *mysqlPortfolioRepository) UpdatePortfolio(idPrtf int, data portfolios.Core) (int, error) {
	var dataPortfolio = fromCore(data)
	result := repo.db.Where("id = ?", idPrtf).Updates(&dataPortfolio)
	if result.Error != nil {
		return 0, result.Error
	}
	return 1, nil
}
