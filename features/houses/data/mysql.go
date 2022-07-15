package data

import (
	"capstoneproject/features/houses"

	"gorm.io/gorm"
)

type mysqlHouseRepository struct {
	db *gorm.DB
}

func NewHouseRepository(conn *gorm.DB) houses.Data {
	return &mysqlHouseRepository{
		db: conn,
	}
}

func (repo *mysqlHouseRepository) SelectAllHouse(limit, offset int) ([]houses.Core, error) {
	var dataHouses []House
	result := repo.db.Preload("User").Limit(limit).Offset(offset).Find(&dataHouses)
	if result.Error != nil {
		return nil, result.Error
	}
	return toCoreList(dataHouses), nil
}
