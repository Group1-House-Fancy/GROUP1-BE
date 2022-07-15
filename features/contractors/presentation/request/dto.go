package request

import "capstoneproject/features/contractors"

type Contractor struct {
	ContractorName      string `json:"contractor_name" form:"contractor_name"`
	NumberSIUJK         string `json:"number_siujk" form:"number_siujk"`
	ImageURL            string `json:"image_url" form:"image_url"`
	CertificateSIUJKURL string `json:"certificate_siujk_url" form:"certificate_siujk_url"`
	PhoneNumber         string `json:"phone_number" form:"phone_number"`
	Email               string `json:"email" form:"email"`
	Address             string `json:"address" form:"address"`
	Description         string `json:"description" form:"description"`
	UserID              uint   `json:"user_id" form:"user_id"`
	UserContractor      bool   `json:"user_contractor" form:"user_contractor"`
}

func ToCore(req Contractor) contractors.Core {
	return contractors.Core{
		ContractorName:      req.ContractorName,
		NumberSIUJK:         req.NumberSIUJK,
		ImageURL:            req.ImageURL,
		CertificateSIUJKURL: req.CertificateSIUJKURL,
		PhoneNumber:         req.PhoneNumber,
		Email:               req.Email,
		Address:             req.Address,
		Description:         req.Description,
		User: contractors.User{
			ID:           int(req.UserID),
			IsContractor: req.UserContractor,
		},
	}
}
