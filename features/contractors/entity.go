package contractors

import "time"

type Core struct {
	ID                  int
	ContractorName      string
	NumberSIUJK         string
	ImageURL            string
	CertificateSIUJKURL string
	PhoneNumber         string
	Email               string
	Address             string
	Description         string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	User                User
}

type User struct {
	ID           int
	IsContractor bool
}

type Business interface {
	CreateContractor(data Core) (row int, err error)
}

type Data interface {
	PostContractor(data Core) (row int, err error)
	ContractorExist(id int, userContractor bool) (row int, err error)
}
