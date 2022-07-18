package response

import (
	"capstoneproject/features/portfolios"
	"strconv"
	"time"
)

type Portfolio struct {
	ID             int                       `json:"id"`
	ClientName     string                    `json:"client_name"`
	Location       string                    `json:"location"`
	FinishDate     string                    `json:"finish_date"`
	Longitude      float64                   `json:"longitude"`
	Latitude       float64                   `json:"latitude"`
	Price          int                       `json:"price"`
	Description    string                    `json:"description"`
	CreatedAt      time.Time                 `json:"created_at"`
	Contractor     Contractor                `json:"contractor"`
	PortfolioImage map[string]PortfolioImage `json:"image_url"`
}

type Contractor struct {
	ID int `json:"id"`
}

type PortfolioImage struct {
	ID       int    `json:"id"`
	ImageURL string `json:"image_url"`
}

func FromCore(data portfolios.Core) Portfolio {
	return Portfolio{
		ID:          data.ID,
		ClientName:  data.ClientName,
		Location:    data.Location,
		FinishDate:  data.FinishDate,
		Longitude:   data.Longitude,
		Latitude:    data.Latitude,
		Price:       data.Price,
		Description: data.Description,
		CreatedAt:   data.CreatedAt,
		Contractor: Contractor{
			ID: data.Contractor.ID,
		},
		PortfolioImage: FromPortfolioImageList(data.PortfolioImage),
	}
}

func FromCoreList(data []portfolios.Core) []Portfolio {
	result := []Portfolio{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}

func FromPortfolioImage(data portfolios.PortfolioImage) PortfolioImage {
	return PortfolioImage{
		ID:       data.ID,
		ImageURL: data.ImageURL,
	}
}

func FromPortfolioImageList(data []portfolios.PortfolioImage) map[string]PortfolioImage {
	result := map[string]PortfolioImage{}
	var index = 1
	for key := range data {
		indexString := strconv.Itoa(index)
		result[indexString] = FromPortfolioImage(data[key])
		index++
	}
	return result
}
