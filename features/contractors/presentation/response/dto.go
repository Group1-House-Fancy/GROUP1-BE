package response

import (
	"capstoneproject/features/contractors"
)

type Contractor struct {
	ID                  int    `json:"id"`
	ContractorName      string `json:"contractor_name"`
	NumberSIUJK         string `json:"number_siujk"`
	ImageURL            string `json:"image_url"`
	PhoneNumber         string `json:"phone_number"`
	Email               string `json:"email"`
	Address             string `json:"address"`
	Description         string `json:"description"`
	CertificateSIUJKURL string `json:"certificate_siujk_url"`
}

type User struct {
	ID           int  `json:"id"`
	IsContractor bool `json:"is_contractor"`
	Contractor   Contractor
}

func FromCore(data contractors.Core) Contractor {
	return Contractor{
		ID:                  data.ID,
		ContractorName:      data.ContractorName,
		NumberSIUJK:         data.NumberSIUJK,
		ImageURL:            data.ImageURL,
		PhoneNumber:         data.PhoneNumber,
		Email:               data.Email,
		Address:             data.Address,
		Description:         data.Description,
		CertificateSIUJKURL: data.CertificateSIUJKURL,
	}
}

func FromCoreList(data []contractors.Core) []Contractor {
	result := []Contractor{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
