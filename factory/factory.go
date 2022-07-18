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
	_negotiationBusiness "capstoneproject/features/negotiations/business"
	_negotiationData "capstoneproject/features/negotiations/data"
	_negotiationPresentation "capstoneproject/features/negotiations/presentation"
	_userBusiness "capstoneproject/features/users/business"
	_userData "capstoneproject/features/users/data"
	_userPresentation "capstoneproject/features/users/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter        *_userPresentation.UserHandler
	ContractorPresenter  *_contractorPresentation.ContractorHandler
	HousePresenter       *_housePresentation.HouseHandler
	HouseImagePresenter  *_houseImagePresentation.HouseImageHandler
	NegotiationPresenter *_negotiationPresentation.NegotiationHandler
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
	negotiationData := _negotiationData.NewNegotiationRepository(dbConn)
	negotiationBusiness := _negotiationBusiness.NewNegotiationBusiness(negotiationData)
	negotiationPresentation := _negotiationPresentation.NewNegotiationHandler(negotiationBusiness)

	return Presenter{
		UserPresenter:        UserPresentation,
		ContractorPresenter:  contractorPresentation,
		HousePresenter:       housePresentation,
		HouseImagePresenter:  houseImagePresentation,
		NegotiationPresenter: negotiationPresentation,
	}
}
