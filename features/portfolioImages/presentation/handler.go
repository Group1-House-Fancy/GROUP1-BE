package presentation

import (
	portfolioimages "capstoneproject/features/portfolioImages"
	_requestPortfolioImage "capstoneproject/features/portfolioImages/presentation/request"
	"capstoneproject/helpers"
	"capstoneproject/middlewares"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PortfolioImageHandler struct {
	portfolioImageBusiness portfolioimages.Business
}

func NewPortfolioImageHandler(business portfolioimages.Business) *PortfolioImageHandler {
	return &PortfolioImageHandler{
		portfolioImageBusiness: business,
	}
}

func (h *PortfolioImageHandler) PostNewPortfolioImage(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helpers.ResponseFailed("invalid token"))
	}
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helpers.ResponseFailed("unauthorized"))
	}
	idPortfolio := c.Param("idPortfolio")
	idPortfolioInt, errIdPortfolio := strconv.Atoi(idPortfolio)
	if errIdPortfolio != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFailed("failed id portfolio not recognize"))
	}
	link, report, errLink := helpers.AddPortfolioImage(c)
	if errLink != nil {
		return c.JSON(report["code"].(int), helpers.ResponseFailed(fmt.Sprintf("%s", report["message"])))
	}
	var dataPortfolioImage = _requestPortfolioImage.PortfolioImage{
		PortfolioID: uint(idPortfolioInt),
		ImageURL:    link,
	}
	result, err := h.portfolioImageBusiness.PostNewPortfolioImage(_requestPortfolioImage.ToCore(dataPortfolioImage))
	if result == -1 {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFailed("all input must be filled"))
	}
	if result == 0 {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to insert image portfolio"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to insert image portfolio"))
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccesNoData("success to insert image portfolio"))
}

func (h *PortfolioImageHandler) DeletePortfolioImage(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helpers.ResponseFailed("invalid token"))
	}
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helpers.ResponseFailed("unauthorized"))
	}
	idImage := c.Param("idPortfolio")
	idImageInt, errIdImage := strconv.Atoi(idImage)
	if errIdImage != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFailed("failed id image not recognize"))
	}
	result, err := h.portfolioImageBusiness.DeleteImage(idImageInt)
	if result == 0 {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to delete image portfolio"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to delete image portfolio"))
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccesNoData("success to delete image portfolio"))
}
