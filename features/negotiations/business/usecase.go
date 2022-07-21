package business

import (
	"capstoneproject/features/negotiations"
	"capstoneproject/plugins"
	"fmt"
)

type negotiationUsecase struct {
	negotiationData negotiations.Data
}

func NewNegotiationBusiness(negoData negotiations.Data) negotiations.Business {
	return &negotiationUsecase{
		negotiationData: negoData,
	}
}

func (uc *negotiationUsecase) GetHistoryUser(idUser, limit, offset int) (resp []negotiations.Core, totalPage int, err error) {
	resp, err = uc.negotiationData.SelectNegotiationsByIdUser(idUser, limit, offset)
	total := len(resp)
	if total == 0 {
		totalPage = 0
	} else {
		if limit == 0 {
			limit = total
		}
		if total%limit != 0 {
			totalPage = (total / limit) + 1
		} else {
			totalPage = total / limit
		}
	}
	return resp, totalPage, err
}

func (uc *negotiationUsecase) GetHouseNegotiators(idHouse, limit, offset int) (resp []negotiations.Core, totalPage int, err error) {
	resp, err = uc.negotiationData.SelectNegotiationsByIdHouse(idHouse, limit, offset)
	total := len(resp)
	if total == 0 {
		totalPage = 0
	} else {
		if limit == 0 {
			limit = total
		}
		if total%limit != 0 {
			totalPage = (total / limit) + 1
		} else {
			totalPage = total / limit
		}
	}
	return resp, totalPage, err
}

func (uc *negotiationUsecase) PostNewNegotiation(input negotiations.Core) (row int, err error) {
	check, _ := uc.negotiationData.CheckAlreadyNegotiation(input.User.ID, input.House.ID)
	if check == 0 {
		if input.Nego == 0 || input.User.ID == 0 || input.House.ID == 0 {
			return -1, fmt.Errorf("all input must be filled")
		}
		row, err = uc.negotiationData.InsertNewNegotiation(input)
		rowHouse, errHouse := uc.negotiationData.UpdateHouseStatus(input.House.ID, "Negotiation")
		if errHouse != nil {
			return rowHouse, errHouse
		}
	} else {
		return -2, fmt.Errorf("already nego")
	}
	return row, err
}

func (uc *negotiationUsecase) UpdateStatus(idNegotiation int, status string) (row int, err error) {
	row, err = uc.negotiationData.UpdateNegotiation(idNegotiation, status)
	data := uc.negotiationData.SelectNegotiation(idNegotiation)
	fmt.Println(data)
	if status == "owned" || status == "Owned" {
		dataNegotiator, _ := uc.negotiationData.SelectNegotiationsByIdHouse(data.House.ID, 0, 0)
		for _, v := range dataNegotiator {
			if v.ID != idNegotiation {
				_, errUpdate := uc.negotiationData.UpdateNegotiation(v.ID, "Cancel")
				if errUpdate != nil {
					return 0, errUpdate
				}
			}
		}
		_, errHouse := uc.negotiationData.UpdateHouseStatus(data.House.ID, "Sold Out")
		if errHouse != nil {
			return 0, errHouse
		}
		result, errUser := uc.negotiationData.SelectUser(data.User.ID)
		if errUser != nil {
			return 0, errUser
		} else {
			go plugins.SendingEmail(result.Email)
		}
	} else if status == "cancel" || status == "Cancel" {
		fmt.Println(uc.negotiationData.CheckNegotiator(data.House.ID))
		if uc.negotiationData.CheckNegotiator(data.House.ID) {
			_, errHouse := uc.negotiationData.UpdateHouseStatus(data.House.ID, "Available")
			if errHouse != nil {
				return 0, errHouse
			}
		}
	}
	return row, err
}

func (uc *negotiationUsecase) DeleteNegotiation(idNegotiation int) (row int, err error) {
	row, err = uc.negotiationData.DeleteNegotiation(idNegotiation)
	return row, err
}
