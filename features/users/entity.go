package users

import "time"

type Core struct {
	ID           int
	FullName     string
	Email        string
	Password     string
	PhoneNumber  string
	Address      string
	ImageURL     string
	IsContractor bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Business interface {
	InsertUser(data Core) (row int, err error)
	LoginUser(email string, password string) (fullName string, token string, imageURL string, isContractor bool, e error)
	UpdateDataUser(id int, data Core) (row int, err error)
	SelectUser(id int) (data Core, err error)
	DeleteUser(id int) (row int, err error)
}

type Data interface {
	PostUser(data Core) (row int, err error)
	AuthUser(email string, password string) (fullName string, token string, imageURL string, isContractor bool, e error)
	PutDataUser(id int, data Core) (row int, err error)
	PutDataUser1(id int, data bool) (row int, err error)
	GetUser(id int) (data Core, err error)
	DeleteUser(id int) (row int, err error)
}
