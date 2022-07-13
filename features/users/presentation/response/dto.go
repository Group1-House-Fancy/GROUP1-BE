package response

import (
	"capstoneproject/features/users"
	"time"
)

type User struct {
	ID           int       `json:"id"`
	FullName     string    `json:"full_name"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phone_number"`
	Address      string    `json:"address"`
	ImageURL     string    `json:"image_url"`
	IsContractor bool      `json:"is_contractor"`
	CreatedAt    time.Time `json:"created_at"`
}

func FromCore(data users.Core) User {
	return User{
		ID:           data.ID,
		FullName:     data.FullName,
		Email:        data.Email,
		PhoneNumber:  data.PhoneNumber,
		Address:      data.Address,
		ImageURL:     data.ImageURL,
		IsContractor: data.IsContractor,
		CreatedAt:    data.CreatedAt,
	}
}

func FromCoreList(data []users.Core) []User {
	result := []User{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
