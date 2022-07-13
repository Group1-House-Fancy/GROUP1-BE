package business

import (
	"capstoneproject/features/users"

	"errors"
)

type userUsecase struct {
	userData users.Data
}

func NewUserBusiness(usrData users.Data) users.Business {
	return &userUsecase{
		userData: usrData,
	}
}

func (uc *userUsecase) InsertUser(input users.Core) (row int, err error) {
	if input.FullName == "" || input.Password == "" || input.Email == "" || input.PhoneNumber == "" || input.Address == "" {
		return -1, errors.New("don't empty for all input")
	}
	row, err = uc.userData.PostUser(input)
	return row, err
}

func (uc *userUsecase) LoginUser(email string, password string) (fullName string, token string, imageURL string, isContractor bool, e error) {
	fullName, token, imageURL, isContractor, e = uc.userData.AuthUser(email, password)
	return fullName, token, imageURL, isContractor, e
}

func (uc *userUsecase) UpdateDataUser(id int, data users.Core) (row int, err error) {
	if data.FullName == "" || data.Password == "" || data.Email == "" || data.PhoneNumber == "" || data.Address == "" || data.ImageURL == "" {
		return -1, errors.New("all input data must be filled")
	}
	row, err = uc.userData.PutDataUser(id, data)
	return row, err
}

func (uc *userUsecase) SelectUser(id int) (resp users.Core, err error) {
	resp, err = uc.userData.GetUser(id)
	return resp, err
}

func (uc *userUsecase) DeleteUser(id int) (row int, err error) {
	row, err = uc.userData.DeleteUser(id)
	return row, err
}
