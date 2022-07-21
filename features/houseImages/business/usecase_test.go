package business

import (
	houseimages "capstoneproject/features/houseImages"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockHouseImageData struct{}

func (mock mockHouseImageData) InsertNewImage(data houseimages.Core) (int, error) {
	return 1, nil
}

func (mock mockHouseImageData) DeleteImage(idImage int) (int, error) {
	return 1, nil
}

type mockHouseImageDataFailed struct{}

func (mock mockHouseImageDataFailed) InsertNewImage(data houseimages.Core) (int, error) {
	return 0, fmt.Errorf("Failed to insert data")
}

func (mock mockHouseImageDataFailed) DeleteImage(idImage int) (int, error) {
	return 0, fmt.Errorf("Failed to delete data")
}

func TestPostNewHouseImage(t *testing.T) {
	t.Run("Test Post New House Image Success", func(t *testing.T) {
		houseImageBusiness := NewHouseImageBusiness(mockHouseImageData{})
		var data = houseimages.Core{
			ImageURL: "storage.cloud.com/event-bersama.jpeg",
			House: houseimages.House{
				ID: 1,
			},
		}
		result, err := houseImageBusiness.PostNewHouseImage(data)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})
	t.Run("Test Post New House Image Failed", func(t *testing.T) {
		houseImageBusiness := NewHouseImageBusiness(mockHouseImageDataFailed{})
		var data = houseimages.Core{
			ImageURL: "storage.cloud.com/event-bersama.jpeg",
			House: houseimages.House{
				ID: 1,
			},
		}
		result, err := houseImageBusiness.PostNewHouseImage(data)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})
	t.Run("Test Post New House Image Failed When URL Empty", func(t *testing.T) {
		houseImageBusiness := NewHouseImageBusiness(mockHouseImageDataFailed{})
		var data = houseimages.Core{
			House: houseimages.House{
				ID: 1,
			},
		}
		result, err := houseImageBusiness.PostNewHouseImage(data)
		assert.NotNil(t, err)
		assert.Equal(t, -1, result)
	})
	t.Run("Test Post New House Image Failed When House ID Empty", func(t *testing.T) {
		houseImageBusiness := NewHouseImageBusiness(mockHouseImageDataFailed{})
		var data = houseimages.Core{
			ImageURL: "storage.cloud.com/event-bersama.jpeg",
		}
		result, err := houseImageBusiness.PostNewHouseImage(data)
		assert.NotNil(t, err)
		assert.Equal(t, -1, result)
	})
}

func TestDeleteImage(t *testing.T) {
	t.Run("Test Delete Image Success", func(t *testing.T) {
		houseImageBusiness := NewHouseImageBusiness(mockHouseImageData{})
		result, err := houseImageBusiness.DeleteImage(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})
	t.Run("Test Delete Image Failed", func(t *testing.T) {
		houseImageBusiness := NewHouseImageBusiness(mockHouseImageDataFailed{})
		result, err := houseImageBusiness.DeleteImage(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})
}
