package data

import (
	"capstoneproject/features/negotiations"

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
	result := repo.db.Preload("User").Preload("House").Limit(limit).Offset(offset).Find(&dataNegotiations)
	if result.Error != nil {
		return nil, result.Error
	}
	return toCoreList(dataNegotiations), nil
}
