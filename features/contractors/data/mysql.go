package data

import (
	"capstoneproject/features/contractors"
	"fmt"

	"gorm.io/gorm"
)

type mysqlContractorRepository struct {
	db *gorm.DB
}

func NewContractorRepository(conn *gorm.DB) contractors.Data {
	return &mysqlContractorRepository{
		db: conn,
	}
}

func (repo *mysqlContractorRepository) PostContractor(input contractors.Core) (int, error) {
	contractor := Contractor{
		ContractorName:      input.ContractorName,
		NumberSIUJK:         input.NumberSIUJK,
		PhoneNumber:         input.PhoneNumber,
		Email:               input.Email,
		Address:             input.Address,
		Description:         input.Description,
		ImageURL:            input.ImageURL,
		CertificateSIUJKURL: input.CertificateSIUJKURL,
		UserID:              uint(input.User.ID),
		UserContractor:      true,
	}
	result := repo.db.Create(&contractor)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlContractorRepository) ContractorExist(id int, bl bool) (int, error) {
	result := repo.db.Where("user_id = ? AND user_contractor = ?", id, true).First(&Contractor{})
	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected != 1 {
		return -1, result.Error
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlContractorRepository) SelectAllContractor(limit, offset int) ([]contractors.Core, error) {
	var dataCr []Contractor
	result := repo.db.Limit(limit).Offset(offset).Find(&dataCr)
	if result.Error != nil {
		return nil, result.Error
	}
	return toCoreList(dataCr), nil
}

func (repo *mysqlContractorRepository) SelectContractor(id int) (contractors.Core, error) {
	var dataCrt Contractor
	result := repo.db.Where("id = ?", id).First(&dataCrt)
	if result.Error != nil {
		return contractors.Core{}, result.Error
	}
	return dataCrt.toCore(), nil
}

func (repo *mysqlContractorRepository) DeleteContractor(id int) (int, error) {
	result := repo.db.Where("user_id = ?", id).Delete(&Contractor{})
	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlContractorRepository) UpdateContractor(idCtr int, idUser int, input contractors.Core) (int, error) {
	var updateCtr = fromCore(input)
	result := repo.db.Where("id = ? AND user_id = ?", idCtr, idUser).Updates(&updateCtr)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, fmt.Errorf("failed to update data")
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlContractorRepository) CountContractorData() (int, error) {
	var count int64
	result := repo.db.Model(&Contractor{}).Count(&count)
	if result.Error != nil {
		return -1, result.Error
	}
	return int(count), nil
}
