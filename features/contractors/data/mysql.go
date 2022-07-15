package data

import (
	"capstoneproject/features/contractors"

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
