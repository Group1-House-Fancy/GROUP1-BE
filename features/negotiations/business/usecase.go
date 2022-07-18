package business

import "capstoneproject/features/negotiations"

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
