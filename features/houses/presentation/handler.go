package presentation

import (
	"capstoneproject/features/houses"
	_requestHouse "capstoneproject/features/houses/presentation/request"
	_responseHouse "capstoneproject/features/houses/presentation/response"
	"capstoneproject/helpers"
	"capstoneproject/middlewares"
	"capstoneproject/plugins"
	"fmt"
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
	if result == 0 {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to insert house"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to insert house"))
	}
	var data = map[string]interface{}{
		"id_house": idHouse,
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccesWithData("Succes to insert house", data))
}

func (h *HouseHandler) GetHouseDetail(c echo.Context) error {
	idHouse := c.Param("idHouse")
	idHouseInt, errIdHouse := strconv.Atoi(idHouse)
	if errIdHouse != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFailed("failed id house not recognize"))
	}
	latSource := c.QueryParam("latitude")
	longSource := c.QueryParam("longitude")
	result, err := h.houseBusiness.GetHouseDetail(idHouseInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to get data"))
	}
	latDest := fmt.Sprintf("%f", result.Latitude)
	longDest := fmt.Sprintf("%f", result.Longitude)
	var resp = map[string]interface{}{
		"data":     _responseHouse.FromCore(result),
		"distance": plugins.DistanceMatrix(latSource, longSource, latDest, longDest),
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccesWithData("success to get data", resp))
}

func (h *HouseHandler) GetMyListHouse(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helpers.ResponseFailed("invalid token"))
	}
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helpers.ResponseFailed("unauthorized"))
	}
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)
	result, totalPage, err := h.houseBusiness.GetMyListHouse(idToken, limitInt, offsetInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to get all data"))
	}
	var resp = map[string]interface{}{
		"data":       _responseHouse.FromCoreList(result),
		"total_page": totalPage,
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccesWithData("success to get all data", resp))
}

func (h *HouseHandler) PutHouse(c echo.Context) error {
	idHouse := c.Param("idHouse")
	idHouseInt, errIdHouse := strconv.Atoi(idHouse)
	if errIdHouse != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFailed("failed id house not recognize"))
	}
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helpers.ResponseFailed("invalid token"))
	}
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helpers.ResponseFailed("unauthorized"))
	}
	var dataHouse = _requestHouse.House{
		UserID: idToken,
	}
	errBind := c.Bind(&dataHouse)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to bind data"))
	}
	result, err := h.houseBusiness.PutHouse(idHouseInt, _requestHouse.ToCore(dataHouse))
	if result == -1 {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFailed("all input must be filled"))
	}
	if result == 0 {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to update house"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to update house"))
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccesNoData("success to update house"))
}

func (h *HouseHandler) DeleteHouse(c echo.Context) error {
	idHouse := c.Param("idHouse")
	idHouseInt, errIdHouse := strconv.Atoi(idHouse)
	if errIdHouse != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFailed("failed id house not recognize"))
	}
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helpers.ResponseFailed("invalid token"))
	}
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helpers.ResponseFailed("unauthorized"))
	}
	result, err := h.houseBusiness.DeleteHouse(idHouseInt)
	if result == 0 {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to delete house"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to delete house"))
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccesNoData("success to delete house"))
}

func (h *HouseHandler) GetSearchHouse(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)
	minPrice := c.QueryParam("min_price")
	maxPrice := c.QueryParam("max_price")
	location := c.QueryParam("location")
	type search struct {
		Keyword string `json:"keyword"`
	}
	var find = search{}
	errBind := c.Bind(&find)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to bind data"))
	}
	result, totalPage, err := h.houseBusiness.GetSearchHouse(find.Keyword, location, minPrice, maxPrice, limitInt, offsetInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to get all data"))
	}
	var resp = map[string]interface{}{
		"data":       _responseHouse.FromCoreList(result),
		"total_page": totalPage,
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccesWithData("success to get all data", resp))
}
