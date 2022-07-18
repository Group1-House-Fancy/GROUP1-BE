package request

import "capstoneproject/features/negotiations"

type Negotiation struct {
	Nego    int    `json:"nego" form:"nego"`
	Status  string `json:"status" form:"status"`
	UserID  uint   `json:"user_id" form:"user_id"`
	HouseID uint   `json:"house_id" form:"house_id"`
}

func ToCore(req Negotiation) negotiations.Core {
	return negotiations.Core{
		Nego:   req.Nego,
		Status: req.Status,
		User: negotiations.User{
			ID: int(req.UserID),
		},
		House: negotiations.House{
			ID: int(req.HouseID),
		},
	}
}
