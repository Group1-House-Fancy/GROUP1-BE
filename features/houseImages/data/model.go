package data

import (
	houseimages "capstoneproject/features/houseImages"

	"gorm.io/gorm"
)

type HouseImage struct {
	gorm.Model
	ImageURL string `json:"image_url" form:"image_url"`
	HouseID  uint   `json:"house_id" form:"house_id"`
}

func (data *HouseImage) toCore() houseimages.Core {
	return houseimages.Core{
		ID:        int(data.ID),
		ImageURL:  data.ImageURL,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		House: houseimages.House{
			ID: int(data.HouseID),
		},
	}
}

func toCoreList(data []HouseImage) []houseimages.Core {
	result := []houseimages.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core houseimages.Core) HouseImage {
	return HouseImage{
		ImageURL: core.ImageURL,
		HouseID:  uint(core.House.ID),
	}
}
