package data

import (
	houseimages "capstoneproject/features/houseImages"

	"gorm.io/gorm"
)

type mysqlHouseImageRepository struct {
	db *gorm.DB
}

func NewHouseImageRepository(conn *gorm.DB) houseimages.Data {
	return &mysqlHouseImageRepository{
		db: conn,
	}
}

func (repo *mysqlHouseImageRepository) InsertNewImages(data houseimages.Core) (int, error) {
	var dataImage = fromCore(data)
	result := repo.db.Create(&dataImage)
	if result.Error != nil {
		return 0, result.Error
	}
	return 1, nil
}
