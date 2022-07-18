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
	e.GET("contractors", presenter.ContractorPresenter.GetAllContractor)
	e.GET("contractors/:idContractor", presenter.ContractorPresenter.GetContractor)
	e.DELETE("contractors", presenter.ContractorPresenter.DeleteContractor, middlewares.JWTMiddleware())
	e.PUT("contractors/:idContractor", presenter.ContractorPresenter.EditContractor, middlewares.JWTMiddleware())

	//houses
	e.GET("houses", presenter.HousePresenter.GetAllHouse)
	e.GET("houses/:idHouse", presenter.HousePresenter.GetHouseDetail)
	e.GET("houses/mylisthouses", presenter.HousePresenter.GetMyListHouse, middlewares.JWTMiddleware())
	e.POST("houses", presenter.HousePresenter.PostNewHouse, middlewares.JWTMiddleware())
	e.PUT("houses/:idHouse", presenter.HousePresenter.PutHouse, middlewares.JWTMiddleware())
	e.DELETE("houses/:idHouse", presenter.HousePresenter.DeleteHouse, middlewares.JWTMiddleware())
	e.POST("houses/images/:idHouse", presenter.HouseImagePresenter.PostNewHouseImage, middlewares.JWTMiddleware())
	e.DELETE("houses/images/:idImage", presenter.HouseImagePresenter.DeleteHouseImage, middlewares.JWTMiddleware())

	//negotiations
	e.GET("negotiations", presenter.NegotiationPresenter.GetHistoryUser, middlewares.JWTMiddleware())
	e.GET("negotiations/:idHouse", presenter.NegotiationPresenter.GetHouseNegotiators, middlewares.JWTMiddleware())

	return e
}
