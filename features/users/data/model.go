package data

import (
	"capstoneproject/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName     string `json:"full_name"`
	Email        string `gorm:"unique" json:"email"`
	Password     string `json:"password"`
	PhoneNumber  string `json:"phone_number"`
	Address      string `json:"address"`
	ImageURL     string `json:"image_url"`
	IsContractor bool   `json:"is_contractor"`
}

func (data *User) toCore() users.Core {
	return users.Core{
		ID:           int(data.ID),
		FullName:     data.FullName,
		Password:     data.Password,
		Email:        data.Email,
		PhoneNumber:  data.PhoneNumber,
		Address:      data.Address,
		ImageURL:     data.ImageURL,
		IsContractor: data.IsContractor,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}

func toCoreList(data []User) []users.Core {
	result := []users.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core users.Core) User {
	return User{
		FullName:     core.FullName,
		Password:     core.Password,
		Email:        core.Email,
		PhoneNumber:  core.PhoneNumber,
		Address:      core.Address,
		ImageURL:     core.ImageURL,
		IsContractor: core.IsContractor,
	}
}
