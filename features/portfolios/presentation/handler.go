package presentation

import (
	"capstoneproject/features/portfolios"
	_requestPortfolio "capstoneproject/features/portfolios/presentation/request"
	_responsePortfolio "capstoneproject/features/portfolios/presentation/response"
	"capstoneproject/helpers"
	"capstoneproject/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PortfolioHandler struct {
	portfolioBusiness portfolios.Business
}

func NewPortfolioHandler(business portfolios.Business) *PortfolioHandler {
	return &PortfolioHandler{
		portfolioBusiness: business,
	}
}

func (h *PortfolioHandler) InsertNewPortfolio(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helpers.ResponseFailed("invalid token"))
	}
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helpers.ResponseFailed("unauthorized"))
	}
	var dataPortfolio _requestPortfolio.Portfolio
	errBind := c.Bind(&dataPortfolio)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to bind data"))
	}
	idPortfolio, result, err := h.portfolioBusiness.PostPortfolio(_requestPortfolio.ToCore(dataPortfolio))
	if result == -1 {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFailed("all input must be filled"))
	}
	if result == 0 {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to insert portfolio"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to insert portfolio"))
	}
	var data = map[string]interface{}{
		"id_portfolio": idPortfolio,
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccesWithData("Succes to insert portfolio", data))
}

func (h *PortfolioHandler) GetAllPortfolio(c echo.Context) error {
	idCtr := c.Param("idContractor")
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	idContractr, errContractr := strconv.Atoi(idCtr)
	if errContractr != nil {
		c.JSON(http.StatusBadRequest, helpers.ResponseFailed("id contractor not recognize"))
	}
	limitNum, _ := strconv.Atoi(limit)
	offsetNum, _ := strconv.Atoi(offset)

	result, totalPage, err := h.portfolioBusiness.GetAllPortfolio(idContractr, limitNum, offsetNum)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to get data"))
	}
	var resp = map[string]interface{}{
		"data":       _responsePortfolio.FromCoreList(result),
		"total_page": totalPage,
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccesWithData("success to get all data ", resp))
}
