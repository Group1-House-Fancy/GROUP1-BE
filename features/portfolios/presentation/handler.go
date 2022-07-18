package presentation

import (
	"capstoneproject/features/portfolios"
	_requestPortfolio "capstoneproject/features/portfolios/presentation/request"
	"capstoneproject/helpers"
	"capstoneproject/middlewares"
	"net/http"

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
