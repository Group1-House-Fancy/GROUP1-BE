package data

import (
	contractors "capstoneproject/features/contractors"

	"gorm.io/gorm"
)

type Contractor struct {
	gorm.Model
	ContractorName      string `json:"contractor_name"`
	NumberSIUJK         string `json:"number_siujk"`
	ImageURL            string `json:"image_url"`
	CertificateSIUJKURL string `json:"certificate_siujk_url"`
	PhoneNumber         string `json:"phone_number"`
	Email               string `json:"email"`
	Address             string `json:"address"`
	Description         string `json:"description"`
	UserID              uint   `json:"user_id"`
	UserContractor      bool   `json:"user_contractor"`
}

type User struct {
	gorm.Model
	IsContractor bool `json:"is_contractor"`
	Contractor   Contractor
}

func (data *Contractor) toCore() contractors.Core {
	return contractors.Core{
		ID:                  int(data.ID),
		ContractorName:      data.ContractorName,
		NumberSIUJK:         data.NumberSIUJK,
		ImageURL:            data.ImageURL,
		CertificateSIUJKURL: data.CertificateSIUJKURL,
		PhoneNumber:         data.PhoneNumber,
		Email:               data.Email,
		Address:             data.Address,
		Description:         data.Description,
		CreatedAt:           data.CreatedAt,
		UpdatedAt:           data.UpdatedAt,
	}
}

func toCoreList(data []Contractor) []contractors.Core {
	result := []contractors.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core contractors.Core) Contractor {
	return Contractor{
		ContractorName:      core.ContractorName,
		NumberSIUJK:         core.NumberSIUJK,
		ImageURL:            core.ImageURL,
		CertificateSIUJKURL: core.CertificateSIUJKURL,
		PhoneNumber:         core.PhoneNumber,
		Email:               core.Email,
		Address:             core.Address,
		Description:         core.Description,
		UserID:              uint(core.User.ID),
	}
}
