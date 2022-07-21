package data

import (
	"capstoneproject/features/negotiations"
	"capstoneproject/features/users"

	"gorm.io/gorm"
)

type mysqlNegotiationRepository struct {
	db *gorm.DB
}

func NewNegotiationRepository(conn *gorm.DB) negotiations.Data {
	return &mysqlNegotiationRepository{
		db: conn,
	}
}

func (repo *mysqlNegotiationRepository) SelectNegotiationsByIdUser(idUser, limit, offset int) ([]negotiations.Core, error) {
	var dataNegotiations = []Negotiation{}
	result := repo.db.Preload("User").Preload("House").Preload("House.HouseImage").Where("user_id = ?", idUser).Limit(limit).Offset(offset).Find(&dataNegotiations)
	if result.Error != nil {
		return nil, result.Error
	}
	return toCoreList(dataNegotiations), nil
}

func (repo *mysqlNegotiationRepository) SelectNegotiationsByIdHouse(idHouse, limit, offset int) ([]negotiations.Core, error) {
	var dataNegotiations = []Negotiation{}
	result := repo.db.Preload("User").Preload("House").Where("house_id = ?", idHouse).Not("status = ?", "Cancel").Order("nego desc").Limit(limit).Offset(offset).Find(&dataNegotiations)
	if result.Error != nil {
		return nil, result.Error
	}
	return toCoreList(dataNegotiations), nil
}

func (repo *mysqlNegotiationRepository) InsertNewNegotiation(data negotiations.Core) (int, error) {
	var dataNegotiation = fromCore(data)
	result := repo.db.Create(&dataNegotiation)
	if result.Error != nil {
		return 0, result.Error
	}
	return 1, nil
}

func (repo *mysqlNegotiationRepository) CheckAlreadyNegotiation(idUser, idHouse int) (row int, err error) {
	result := repo.db.Where("user_id = ? AND house_id = ? ", idUser, idHouse).Not("status = ?", "Cancel").First(&Negotiation{})
	if result.RowsAffected == 0 {
		return 0, nil
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlNegotiationRepository) UpdateHouseStatus(idHouse int, status string) (int, error) {
	result := repo.db.Model(&House{}).Where("id = ? ", idHouse).Update("status", status)
	if result.Error != nil {
		return 0, result.Error
	}
	return 1, nil
}

func (repo *mysqlNegotiationRepository) UpdateNegotiation(idNegotiation int, status string) (int, error) {
	result := repo.db.Model(&Negotiation{}).Where("id = ? ", idNegotiation).Update("status", status)
	if result.Error != nil {
		return 0, result.Error
	}
	return 1, nil
}

func (repo *mysqlNegotiationRepository) SelectNegotiation(idNegotiation int) negotiations.Core {
	var dataNegotiation = Negotiation{}
	result := repo.db.Where("id = ?", idNegotiation).First(&dataNegotiation)
	if result.Error != nil {
		return negotiations.Core{}
	}
	return dataNegotiation.toCore()
}

func (repo *mysqlNegotiationRepository) CheckNegotiator(idHouse int) bool {
	result := repo.db.Where("house_id = ? ", idHouse).Not("status = ?", "Cancel").First(&Negotiation{})
	return result.RowsAffected == 0
}

func (repo *mysqlNegotiationRepository) DeleteNegotiation(idNegotiation int) (int, error) {
	result := repo.db.Where("id = ?", idNegotiation).Delete(&Negotiation{})
	if result.Error != nil {
		return 0, result.Error
	}
	return 1, nil
}

func (repo *mysqlNegotiationRepository) SelectUser(idUser int) (users.Core, error) {
	var dataEmail = User{}
	result := repo.db.Where("id = ? ", idUser).First(&dataEmail)
	if result.Error != nil {
		return users.Core{}, result.Error
	}
	return dataEmail.toCoreUser(), nil
}
