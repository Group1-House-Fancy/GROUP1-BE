package factory

import (
	_contractorBusiness "capstoneproject/features/contractors/business"
	_contractorData "capstoneproject/features/contractors/data"
	_contractorPresentation "capstoneproject/features/contractors/presentation"
	_houseImageBusiness "capstoneproject/features/houseImages/business"
	_houseImageData "capstoneproject/features/houseImages/data"
	_houseImagePresentation "capstoneproject/features/houseImages/presentation"
	_houseBusiness "capstoneproject/features/houses/business"
	_houseData "capstoneproject/features/houses/data"
	_housePresentation "capstoneproject/features/houses/presentation"
	_portfolioImageBusiness "capstoneproject/features/portfolioImages/business"
	_portfolioImageData "capstoneproject/features/portfolioImages/data"
	_portfolioImagePresentation "capstoneproject/features/portfolioImages/presentation"
	_portfolioBusiness "capstoneproject/features/portfolios/business"
	_portfolioData "capstoneproject/features/portfolios/data"
	_portfolioPresentation "capstoneproject/features/portfolios/presentation"
	_userBusiness "capstoneproject/features/users/business"
	_userData "capstoneproject/features/users/data"
	_userPresentation "capstoneproject/features/users/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter           *_userPresentation.UserHandler
	ContractorPresenter     *_contractorPresentation.ContractorHandler
	HousePresenter          *_housePresentation.HouseHandler
	HouseImagePresenter     *_houseImagePresentation.HouseImageHandler
	PortfolioPresenter      *_portfolioPresentation.PortfolioHandler
	PortfolioImagePresenter *_portfolioImagePresentation.PortfolioImageHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	UserPresentation := _userPresentation.NewUserHandler(userBusiness)
	contactorData := _contractorData.NewContractorRepository(dbConn)
	contractorBusiness := _contractorBusiness.NewContractorBusiness(contactorData, userData)
	contractorPresentation := _contractorPresentation.NewContractorHandler(contractorBusiness)
	houseData := _houseData.NewHouseRepository(dbConn)
	houseBusiness := _houseBusiness.NewHouseBusiness(houseData)
	housePresentation := _housePresentation.NewHouseHandler(houseBusiness)
	houseImageData := _houseImageData.NewHouseImageRepository(dbConn)
	houseImageBusiness := _houseImageBusiness.NewHouseImageBusiness(houseImageData)
	houseImagePresentation := _houseImagePresentation.NewHouseImageHandler(houseImageBusiness)
	portfolioData := _portfolioData.NewPortfolioRepository(dbConn)
	portfolioBusiness := _portfolioBusiness.NewPortfolioBusiness(portfolioData)
	portfolioPresentation := _portfolioPresentation.NewPortfolioHandler(portfolioBusiness)
	portfolioImageData := _portfolioImageData.NewPortfolioImageRepository(dbConn)
	portfolioImageBusiness := _portfolioImageBusiness.NewPortfolioImageBusiness(portfolioImageData)
	portfolioImagePresentation := _portfolioImagePresentation.NewPortfolioImageHandler(portfolioImageBusiness)

	return Presenter{
		UserPresenter:           UserPresentation,
		ContractorPresenter:     contractorPresentation,
		HousePresenter:          housePresentation,
		HouseImagePresenter:     houseImagePresentation,
		PortfolioPresenter:      portfolioPresentation,
		PortfolioImagePresenter: portfolioImagePresentation,
	}
}
