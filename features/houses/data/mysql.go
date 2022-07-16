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

func (repo *mysqlHouseRepository) InsertNewHouse(data houses.Core) (int, int, error) {
	var dataHouse = fromCore(data)
	result := repo.db.Create(&dataHouse)
	if result.Error != nil {
		return 0, 0, result.Error
	}
	return int(dataHouse.ID), 1, nil
}

func (repo *mysqlHouseRepository) SelectHouseByIdHouse(idHouse int) (houses.Core, error) {
	var dataHouses House
	result := repo.db.Preload("User").Where("id = ?", idHouse).First(&dataHouses)
	if result.Error != nil {
		return houses.Core{}, result.Error
	}
	return dataHouses.toCore(), nil
}

func (repo *mysqlHouseRepository) SelectHouseByIdUser(idUser, limit, offset int) ([]houses.Core, error) {
	var dataHouses []House
	result := repo.db.Preload("User").Where("user_id = ?", idUser).Limit(limit).Offset(offset).Find(&dataHouses)
	if result.Error != nil {
		return nil, result.Error
	}
	return toCoreList(dataHouses), nil
}

func (repo *mysqlHouseRepository) UpdateHouse(idHouse int, data houses.Core) (int, error) {
	var dataHouse = fromCore(data)
	result := repo.db.Where("id = ?", idHouse).Updates(&dataHouse)
	if result.Error != nil {
		return 0, result.Error
	}
	return 1, nil
}
func (repo *mysqlHouseRepository) DeleteHouse(idHouse int) (int, error) {
	result := repo.db.Where("id = ?", idHouse).Delete(&House{})
	if result.Error != nil {
		return 0, result.Error
	}
	return 1, nil
}
