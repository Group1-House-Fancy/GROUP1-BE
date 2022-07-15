package factory

import (
	_houseBusiness "capstoneproject/features/houses/business"
	_houseData "capstoneproject/features/houses/data"
	_housePresentation "capstoneproject/features/houses/presentation"
	_userBusiness "capstoneproject/features/users/business"
	_userData "capstoneproject/features/users/data"
	_userPresentation "capstoneproject/features/users/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter  *_userPresentation.UserHandler
	HousePresenter *_housePresentation.HouseHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	UserPresentation := _userPresentation.NewUserHandler(userBusiness)

	houseData := _houseData.NewHouseRepository(dbConn)
	houseBusiness := _houseBusiness.NewHouseBusiness(houseData)
	housePresentation := _housePresentation.NewHouseHandler(houseBusiness)

	return Presenter{
		UserPresenter:  UserPresentation,
		HousePresenter: housePresentation,
	}
}
