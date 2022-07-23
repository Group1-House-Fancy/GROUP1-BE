package business

import (
	"capstoneproject/features/portfolioImages"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockPortfolioImageData struct{}

func (mock mockPortfolioImageData) InsertNewImage(data portfolioImages.Core) (int, error) {
	return 1, nil
}

func (mock mockPortfolioImageData) DeleteImage(idImage int) (int, error) {
	return 1, nil
}

type mockPortfolioImageDataFailed struct{}

func (mock mockPortfolioImageDataFailed) InsertNewImage(data portfolioImages.Core) (int, error) {
	return 0, fmt.Errorf("failed to insert data")
}

func (mock mockPortfolioImageDataFailed) DeleteImage(idImage int) (int, error) {
	return 0, fmt.Errorf("failed to delete data")
}

func TestPostNewPortfolioImage(t *testing.T) {
	t.Run("Test Post New Portfolio Image Success", func(t *testing.T) {
		portfolioImageBusiness := NewPortfolioImageBusiness(mockPortfolioImageData{})
		var data = portfolioImages.Core{
			ImageURL: "storage.cloud.com/portfolio.jpeg",
			Portfolio: portfolioImages.Portfolio{
				ID: 1,
			},
		}
		result, err := portfolioImageBusiness.PostNewPortfolioImage(data)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})
	t.Run("Test Post New Portfolio Image Failed", func(t *testing.T) {
		portfolioImageBusiness := NewPortfolioImageBusiness(mockPortfolioImageDataFailed{})
		var data = portfolioImages.Core{
			ImageURL: "storage.cloud.com/portfolio.jpeg",
			Portfolio: portfolioImages.Portfolio{
				ID: 1,
			},
		}
		result, err := portfolioImageBusiness.PostNewPortfolioImage(data)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})
	t.Run("Test Post New Portfolio Image Failed When URL Empty", func(t *testing.T) {
		portfolioImageBusiness := NewPortfolioImageBusiness(mockPortfolioImageDataFailed{})
		var data = portfolioImages.Core{
			Portfolio: portfolioImages.Portfolio{
				ID: 1,
			},
		}
		result, err := portfolioImageBusiness.PostNewPortfolioImage(data)
		assert.NotNil(t, err)
		assert.Equal(t, -1, result)
	})
	t.Run("Test Post New Portfolio Image Failed When Portfolio ID empty", func(t *testing.T) {
		portfolioImageBusiness := NewPortfolioImageBusiness(mockPortfolioImageDataFailed{})
		var data = portfolioImages.Core{
			ImageURL: "storage.cloud.com/portfolio.jpeg",
		}
		result, err := portfolioImageBusiness.PostNewPortfolioImage(data)
		assert.NotNil(t, err)
		assert.Equal(t, -1, result)
	})
}

func TestDeleteImage(t *testing.T) {
	t.Run("Test Delete Image Success", func(t *testing.T) {
		portfolioImageBusiness := NewPortfolioImageBusiness(mockPortfolioImageData{})
		result, err := portfolioImageBusiness.DeleteImage(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})
	t.Run("Test Delete Image Failed", func(t *testing.T) {
		portfolioImageBusiness := NewPortfolioImageBusiness(mockPortfolioImageDataFailed{})
		result, err := portfolioImageBusiness.DeleteImage(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})
}
