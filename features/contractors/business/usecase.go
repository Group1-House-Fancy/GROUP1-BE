package business

import (
	"capstoneproject/features/contractors"
	"capstoneproject/features/users"
	"errors"
)

type contractorUsecase struct {
	contractorData contractors.Data
	userData       users.Data
}

func NewContractorBusiness(crData contractors.Data, usrData users.Data) contractors.Business {
	return &contractorUsecase{
		contractorData: crData,
		userData:       usrData,
	}
}

func (uc *contractorUsecase) CreateContractor(input contractors.Core) (row int, err error) {
	if input.ContractorName == "" || input.NumberSIUJK == "" || input.PhoneNumber == "" || input.Address == "" || input.Description == "" || input.ImageURL == "" || input.CertificateSIUJKURL == "" {
		return -1, errors.New("all input data must be filled")
	}
	check, _ := uc.contractorData.ContractorExist(input.User.ID, input.User.IsContractor)
	if check == 1 {
		return -1, errors.New("already as join contractor")
	} else {
		row, err = uc.contractorData.PostContractor(input)
		if row != 0 {
			row1, err1 := uc.userData.PutDataUser1(input.User.ID, input.User.IsContractor)
			return row1, err1
		}
		return row, err
	}
}

func (uc *contractorUsecase) GetAllContractor(limit, offset int) (data []contractors.Core, totalPage int, err error) {
	data, err = uc.contractorData.SelectAllContractor(limit, offset)
	total := len(data)
	if total == 0 {
		totalPage = 0
	} else {
		if limit == 0 {
			limit = total
		}
		if total%limit != 0 {
			totalPage = (total / limit) + 1
		} else {
			totalPage = total / limit
		}
	}
	return data, totalPage, err
}

func (uc *contractorUsecase) GetContractor(id int) (data contractors.Core, err error) {
	data, err = uc.contractorData.SelectContractor(id)
	return data, err
}

func (uc *contractorUsecase) DeleteContractor(idUser int) (row int, err error) {
	row, err = uc.contractorData.DeleteContractor(idUser)
	return row, err
}

func (uc *contractorUsecase) PutContractor(idCtr int, idUser int, data contractors.Core) (row int, err error) {
	if data.ContractorName == "" || data.NumberSIUJK == "" || data.PhoneNumber == "" || data.Address == "" || data.Description == "" || data.ImageURL == "" || data.CertificateSIUJKURL == "" {
		return -1, errors.New("all input data must be filled")
	}
	row, err = uc.contractorData.UpdateContractor(idCtr, idUser, data)
	return row, err
}
