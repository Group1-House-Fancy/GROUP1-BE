package presentation

import (
	"capstoneproject/features/houses"
	_responseHouse "capstoneproject/features/houses/presentation/response"
	"capstoneproject/helpers"
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
