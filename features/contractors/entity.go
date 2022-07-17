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
	GetAllContractor(limit, offset int) (data []Core, totalPage int, err error)
	GetContractor(idCtr int) (data Core, err error)
	DeleteContractor(idUser int) (row int, err error)
	PutContractor(idCtr int, idUser int, data Core) (row int, err error)
}

type Data interface {
	PostContractor(data Core) (row int, err error)
	ContractorExist(id int, userContractor bool) (row int, err error)
	SelectAllContractor(limit, offset int) (data []Core, err error)
	SelectContractor(idCtr int) (data Core, err error)
	DeleteContractor(idUser int) (row int, err error)
	UpdateContractor(idCtr int, idUser int, data Core) (row int, err error)
}
