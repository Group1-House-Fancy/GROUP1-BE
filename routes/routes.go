package routes

import (
	"capstoneproject/factory"
	"capstoneproject/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(presenter factory.Presenter) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.Pre(middleware.RemoveTrailingSlash())
	e.POST("users", presenter.UserPresenter.AddUser)
	e.POST("login", presenter.UserPresenter.Login)
	e.PUT("users", presenter.UserPresenter.EditData, middlewares.JWTMiddleware())
	e.GET("users", presenter.UserPresenter.GetUser, middlewares.JWTMiddleware())
	e.DELETE("users", presenter.UserPresenter.DeleteDataUser, middlewares.JWTMiddleware())

	//contractor
	e.POST("contractors", presenter.ContractorPresenter.JoinContractor, middlewares.JWTMiddleware())

	//houses
	e.GET("houses", presenter.HousePresenter.GetAllHouse)

	return e
}
