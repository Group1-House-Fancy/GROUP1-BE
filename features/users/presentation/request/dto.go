package request

import "capstoneproject/features/users"

type User struct {
	FullName    string `json:"full_name" form:"full_name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Address     string `json:"address" form:"address"`
	ImageURL    string `json:"image_url" form:"image_url"`
}

func ToCore(req User) users.Core {
	return users.Core{
		FullName:    req.FullName,
		Password:    req.Password,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		ImageURL:    req.ImageURL,
	}
}
