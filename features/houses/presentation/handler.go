package presentation

import (
	"capstoneproject/features/houses"
	_requestHouse "capstoneproject/features/houses/presentation/request"
	_responseHouse "capstoneproject/features/houses/presentation/response"
	"capstoneproject/helpers"
	"capstoneproject/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HouseHandler struct {
	houseBusiness houses.Business
}

func NewHouseHandler(business houses.Business) *HouseHandler {
	return &HouseHandler{
		houseBusiness: business,
	}
}

func (h *HouseHandler) GetAllHouse(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)
	result, totalPage, err := h.houseBusiness.GetAllHouse(limitInt, offsetInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to get all data"))
	}
	var resp = map[string]interface{}{
		"data":       _responseHouse.FromCoreList(result),
		"total_page": totalPage,
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccesWithData("success to get all data", resp))
}

func (h *HouseHandler) PostNewHouse(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helpers.ResponseFailed("invalid token"))
	}
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helpers.ResponseFailed("unauthorized"))
	}
	var dataHouse = _requestHouse.House{
		Status: "Available",
		UserID: idToken,
	}
	errBind := c.Bind(&dataHouse)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to bind data"))
	}
	idHouse, result, err := h.houseBusiness.PostNewHouse(_requestHouse.ToCore(dataHouse))
	if result == -1 {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFailed("all input must be filled"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to insert house"))
	}
	var data = map[string]interface{}{
		"id_house": idHouse,
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccesWithData("Succes to insert house", data))
}
