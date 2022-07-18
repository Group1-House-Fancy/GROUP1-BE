package presentation

import (
	"capstoneproject/features/negotiations"
	_responseNegotiation "capstoneproject/features/negotiations/presentation/response"
	"capstoneproject/helpers"
	"capstoneproject/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type NegotiationHandler struct {
	NegotiationBusiness negotiations.Business
}

func NewNegotiationHandler(business negotiations.Business) *NegotiationHandler {
	return &NegotiationHandler{
		NegotiationBusiness: business,
	}
}

func (h *NegotiationHandler) GetHistoryUser(c echo.Context) error {
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
	result, totalPage, err := h.NegotiationBusiness.GetHistoryUser(idToken, limitInt, offsetInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to get all data"))
	}
	var resp = map[string]interface{}{
		"data":       _responseNegotiation.FromCoreHistoryList(result),
		"total_page": totalPage,
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccesWithData("success to get all data", resp))
}

func (h *NegotiationHandler) GetHouseNegotiators(c echo.Context) error {
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
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)
	result, totalPage, err := h.NegotiationBusiness.GetHouseNegotiators(idHouseInt, limitInt, offsetInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to get all data"))
	}
	var resp = map[string]interface{}{
		"data":       _responseNegotiation.FromCoreNegotiatorList(result),
		"total_page": totalPage,
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccesWithData("success to get all data", resp))
}
