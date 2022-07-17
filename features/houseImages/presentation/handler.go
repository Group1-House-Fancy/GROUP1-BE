package presentation

import (
	houseimages "capstoneproject/features/houseImages"
	_requestHouseImage "capstoneproject/features/houseImages/presentation/request"
	"capstoneproject/helpers"
	"capstoneproject/middlewares"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HouseImageHandler struct {
	houseImageBusiness houseimages.Business
}

func NewHouseImageHandler(business houseimages.Business) *HouseImageHandler {
	return &HouseImageHandler{
		houseImageBusiness: business,
	}
}

func (h *HouseImageHandler) PostNewHouseImage(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helpers.ResponseFailed("invalid token"))
	}
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helpers.ResponseFailed("unauthorized"))
	}
	idHouse := c.Param("idHouse")
	idHouseInt, errIdHouse := strconv.Atoi(idHouse)
	if errIdHouse != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFailed("failed id house not recognize"))
	}
	link, report, errLink := helpers.AddHouseImage(c)
	if errLink != nil {
		return c.JSON(report["code"].(int), helpers.ResponseFailed(fmt.Sprintf("%s", report["message"])))
	}
	var dataHouseImage = _requestHouseImage.HouseImage{
		HouseID:  uint(idHouseInt),
		ImageURL: link,
	}
	result, err := h.houseImageBusiness.PostNewHouseImage(_requestHouseImage.ToCore(dataHouseImage))
	if result == -1 {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFailed("all input must be filled"))
	}
	if result == 0 {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to insert image house"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to insert image house"))
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccesNoData("success to insert image house"))
}

func (h *HouseImageHandler) DeleteHouseImage(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helpers.ResponseFailed("invalid token"))
	}
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helpers.ResponseFailed("unauthorized"))
	}
	idImage := c.Param("idImage")
	idImageInt, errIdHouse := strconv.Atoi(idImage)
	if errIdHouse != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFailed("failed id house not recognize"))
	}
	result, err := h.houseImageBusiness.DeleteImage(idImageInt)
	if result == 0 {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to delete image house"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to delete image house"))
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccesNoData("success to delete image house"))
}
