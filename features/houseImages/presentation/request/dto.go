package request

import houseimages "capstoneproject/features/houseImages"

type HouseImage struct {
	ImageURL string `json:"image_url" form:"image_url"`
	HouseID  uint   `json:"house_id" form:"house_id"`
}

func ToCore(req HouseImage) houseimages.Core {
	return houseimages.Core{
		ImageURL: req.ImageURL,
		House: houseimages.House{
			ID: int(req.HouseID),
		},
	}
}
