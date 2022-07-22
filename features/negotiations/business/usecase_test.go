package business

import (
	"capstoneproject/features/negotiations"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockNegotiationData struct{}

func (mock mockNegotiationData) SelectNegotiationsByIdUser(idUser, limit, offset int) ([]negotiations.Core, error) {
	return []negotiations.Core{
		{
			ID:     1,
			Nego:   250000000,
			Status: "Negotiation",
			User: negotiations.User{
				ID: 1,
			},
			House: negotiations.House{
				ID: 1,
			},
		}, {
			ID:     2,
			Nego:   270000000,
			Status: "Cancel",
			User: negotiations.User{
				ID: 1,
			},
			House: negotiations.House{
				ID: 2,
			},
		}, {
			ID:     3,
			Nego:   300000000,
			Status: "Owned",
			User: negotiations.User{
				ID: 1,
			},
			House: negotiations.House{
				ID: 3,
			},
		},
	}, nil
}

func (mock mockNegotiationData) SelectNegotiationsByIdHouse(idHouse, limit, offset int) ([]negotiations.Core, error) {
	return []negotiations.Core{
		{
			ID:     1,
			Nego:   250000000,
			Status: "Negotiation",
			User: negotiations.User{
				ID: 1,
			},
			House: negotiations.House{
				ID: 3,
			},
		}, {
			ID:     2,
			Nego:   300000000,
			Status: "Negotiation",
			User: negotiations.User{
				ID: 2,
			},
			House: negotiations.House{
				ID: 3,
			},
		},
	}, nil
}

func (mock mockNegotiationData) InsertNewNegotiation(data negotiations.Core) (int, error) {
	return 1, nil
}

func (mock mockNegotiationData) CheckAlreadyNegotiation(idUser, idHouse int) (int, error) {
	return 0, nil
}

func (mock mockNegotiationData) UpdateHouseStatus(idHouse int, status string) (int, error) {
	return 1, nil
}

func (mock mockNegotiationData) UpdateNegotiation(idNegotiation int, status string) (int, error) {
	return 1, nil
}

func (mock mockNegotiationData) SelectNegotiation(idNegotiation int) negotiations.Core {
	return negotiations.Core{
		ID:     1,
		Nego:   250000000,
		Status: "Owned",
		User:   negotiations.User{ID: 1},
		House:  negotiations.House{ID: 1},
	}
}

func (mock mockNegotiationData) CheckNegotiator(idHouse int) bool {
	return true
}

func (mock mockNegotiationData) DeleteNegotiation(idNegotiation int) (int, error) {
	return 1, nil
}

func (mock mockNegotiationData) CountHistoryData(idUser int) (int, error) {
	return 5, nil
}

func (mock mockNegotiationData) CountNegotiatorData(idHouse int) (int, error) {
	return 5, nil
}

type mockNegotiationDataFailed struct{}

func (mock mockNegotiationDataFailed) SelectNegotiationsByIdUser(idUser, limit, offset int) ([]negotiations.Core, error) {
	return nil, fmt.Errorf("error to get all data")
}

func (mock mockNegotiationDataFailed) SelectNegotiationsByIdHouse(idHouse, limit, offset int) ([]negotiations.Core, error) {
	return nil, fmt.Errorf("error to get all data")
}

func (mock mockNegotiationDataFailed) InsertNewNegotiation(data negotiations.Core) (int, error) {
	return 0, fmt.Errorf("error to insert data")
}

func (mock mockNegotiationDataFailed) CheckAlreadyNegotiation(idUser, idHouse int) (int, error) {
	return 1, nil
}

func (mock mockNegotiationDataFailed) UpdateHouseStatus(idHouse int, status string) (int, error) {
	return 0, fmt.Errorf("error to update house status")
}

func (mock mockNegotiationDataFailed) UpdateNegotiation(idNegotiation int, status string) (int, error) {
	return 0, fmt.Errorf("error to update data")
}

func (mock mockNegotiationDataFailed) SelectNegotiation(idNegotiation int) negotiations.Core {
	return negotiations.Core{}
}

func (mock mockNegotiationDataFailed) CheckNegotiator(idHouse int) bool {
	return false
}

func (mock mockNegotiationDataFailed) DeleteNegotiation(idNegotiation int) (int, error) {
	return 0, fmt.Errorf("error to delete data")
}

func (mock mockNegotiationDataFailed) CountHistoryData(idUser int) (int, error) {
	return 0, fmt.Errorf("error to count data")
}

func (mock mockNegotiationDataFailed) CountNegotiatorData(idHouse int) (int, error) {
	return 0, fmt.Errorf("error to count data")
}

func TestGetHistoryUser(t *testing.T) {
	t.Run("Test Get History User Success", func(t *testing.T) {
		negotiationBusiness := NewNegotiationBusiness(mockNegotiationData{})
		result, totalPage, err := negotiationBusiness.GetHistoryUser(1, 0, 0)
		assert.Nil(t, err)
		assert.Equal(t, []negotiations.Core{
			{
				ID:     1,
				Nego:   250000000,
				Status: "Negotiation",
				User: negotiations.User{
					ID: 1,
				},
				House: negotiations.House{
					ID: 1,
				},
			}, {
				ID:     2,
				Nego:   270000000,
				Status: "Cancel",
				User: negotiations.User{
					ID: 1,
				},
				House: negotiations.House{
					ID: 2,
				},
			}, {
				ID:     3,
				Nego:   300000000,
				Status: "Owned",
				User: negotiations.User{
					ID: 1,
				},
				House: negotiations.House{
					ID: 3,
				},
			},
		}, result)
		assert.Equal(t, 1, totalPage)
	})
	t.Run("Test Get History User Success When Limit is Odd", func(t *testing.T) {
		negotiationBusiness := NewNegotiationBusiness(mockNegotiationData{})
		result, totalPage, err := negotiationBusiness.GetHistoryUser(1, 1, 0)
		assert.Nil(t, err)
		assert.Equal(t, []negotiations.Core{
			{
				ID:     1,
				Nego:   250000000,
				Status: "Negotiation",
				User: negotiations.User{
					ID: 1,
				},
				House: negotiations.House{
					ID: 1,
				},
			}, {
				ID:     2,
				Nego:   270000000,
				Status: "Cancel",
				User: negotiations.User{
					ID: 1,
				},
				House: negotiations.House{
					ID: 2,
				},
			}, {
				ID:     3,
				Nego:   300000000,
				Status: "Owned",
				User: negotiations.User{
					ID: 1,
				},
				House: negotiations.House{
					ID: 3,
				},
			},
		}, result)
		assert.Equal(t, 5, totalPage)
	})
	t.Run("Test Get History User Success When Limit is Even", func(t *testing.T) {
		negotiationBusiness := NewNegotiationBusiness(mockNegotiationData{})
		result, totalPage, err := negotiationBusiness.GetHistoryUser(1, 2, 0)
		assert.Nil(t, err)
		assert.Equal(t, []negotiations.Core{
			{
				ID:     1,
				Nego:   250000000,
				Status: "Negotiation",
				User: negotiations.User{
					ID: 1,
				},
				House: negotiations.House{
					ID: 1,
				},
			}, {
				ID:     2,
				Nego:   270000000,
				Status: "Cancel",
				User: negotiations.User{
					ID: 1,
				},
				House: negotiations.House{
					ID: 2,
				},
			}, {
				ID:     3,
				Nego:   300000000,
				Status: "Owned",
				User: negotiations.User{
					ID: 1,
				},
				House: negotiations.House{
					ID: 3,
				},
			},
		}, result)
		assert.Equal(t, 3, totalPage)
	})
	t.Run("Test Get History User Failed", func(t *testing.T) {
		negotiationBusiness := NewNegotiationBusiness(mockNegotiationDataFailed{})
		result, totalPage, err := negotiationBusiness.GetHistoryUser(1, 0, 0)
		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Equal(t, 0, totalPage)
	})
}

func TestGetHouseNegotiators(t *testing.T) {
	t.Run("Test Get House Negotiators Success", func(t *testing.T) {
		negotiationBusiness := NewNegotiationBusiness(mockNegotiationData{})
		result, totalPage, err := negotiationBusiness.GetHouseNegotiators(3, 0, 0)
		assert.Nil(t, err)
		assert.Equal(t, []negotiations.Core{
			{
				ID:     1,
				Nego:   250000000,
				Status: "Negotiation",
				User: negotiations.User{
					ID: 1,
				},
				House: negotiations.House{
					ID: 3,
				},
			}, {
				ID:     2,
				Nego:   300000000,
				Status: "Negotiation",
				User: negotiations.User{
					ID: 2,
				},
				House: negotiations.House{
					ID: 3,
				},
			},
		}, result)
		assert.Equal(t, 1, totalPage)
	})
	t.Run("Test Get House Negotiators Success When Limit is Odd", func(t *testing.T) {
		negotiationBusiness := NewNegotiationBusiness(mockNegotiationData{})
		result, totalPage, err := negotiationBusiness.GetHouseNegotiators(3, 1, 0)
		assert.Nil(t, err)
		assert.Equal(t, []negotiations.Core{
			{
				ID:     1,
				Nego:   250000000,
				Status: "Negotiation",
				User: negotiations.User{
					ID: 1,
				},
				House: negotiations.House{
					ID: 3,
				},
			}, {
				ID:     2,
				Nego:   300000000,
				Status: "Negotiation",
				User: negotiations.User{
					ID: 2,
				},
				House: negotiations.House{
					ID: 3,
				},
			},
		}, result)
		assert.Equal(t, 5, totalPage)
	})
	t.Run("Test Get House Negotiators Success When Limit is Even", func(t *testing.T) {
		negotiationBusiness := NewNegotiationBusiness(mockNegotiationData{})
		result, totalPage, err := negotiationBusiness.GetHouseNegotiators(3, 2, 0)
		assert.Nil(t, err)
		assert.Equal(t, []negotiations.Core{
			{
				ID:     1,
				Nego:   250000000,
				Status: "Negotiation",
				User: negotiations.User{
					ID: 1,
				},
				House: negotiations.House{
					ID: 3,
				},
			}, {
				ID:     2,
				Nego:   300000000,
				Status: "Negotiation",
				User: negotiations.User{
					ID: 2,
				},
				House: negotiations.House{
					ID: 3,
				},
			},
		}, result)
		assert.Equal(t, 3, totalPage)
	})
	t.Run("Test Get House Negotiators Failed", func(t *testing.T) {
		negotiationBusiness := NewNegotiationBusiness(mockNegotiationDataFailed{})
		result, totalPage, err := negotiationBusiness.GetHouseNegotiators(3, 0, 0)
		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Equal(t, 0, totalPage)
	})
}

func TestPostNewNegotiation(t *testing.T) {
	t.Run("Test Post New Negotiation Success", func(t *testing.T) {
		negotiationBusiness := NewNegotiationBusiness(mockNegotiationData{})
		var data = negotiations.Core{
			Status: "Negotiation",
			Nego:   250000000,
			User: negotiations.User{
				ID: 1,
			},
			House: negotiations.House{
				ID: 1,
			},
		}
		row, err := negotiationBusiness.PostNewNegotiation(data)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})
	t.Run("Test Post New Negotiation Failed", func(t *testing.T) {
		negotiationBusiness := NewNegotiationBusiness(mockNegotiationDataFailed{})
		var data = negotiations.Core{
			Status: "Negotiation",
			Nego:   250000000,
			User: negotiations.User{
				ID: 1,
			},
			House: negotiations.House{
				ID: 1,
			},
		}
		row, err := negotiationBusiness.PostNewNegotiation(data)
		assert.NotNil(t, err)
		assert.Equal(t, -2, row)
	})
	t.Run("Test Post New Negotiation Failed When Nego is Empty", func(t *testing.T) {
		negotiationBusiness := NewNegotiationBusiness(mockNegotiationData{})
		var data = negotiations.Core{
			Status: "Negotiation",
			User: negotiations.User{
				ID: 1,
			},
			House: negotiations.House{
				ID: 1,
			},
		}
		row, err := negotiationBusiness.PostNewNegotiation(data)
		assert.NotNil(t, err)
		assert.Equal(t, -1, row)
	})
}

func TestUpdateStatus(t *testing.T) {
	t.Run("Test Update Status Success When Status is Owned", func(t *testing.T) {
		negotiationBusiness := NewNegotiationBusiness(mockNegotiationData{})
		row, err := negotiationBusiness.UpdateStatus(1, "owned")
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})
	t.Run("Test Update Status Success When Status is Cancel", func(t *testing.T) {
		negotiationBusiness := NewNegotiationBusiness(mockNegotiationData{})
		row, err := negotiationBusiness.UpdateStatus(1, "cancel")
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})
	t.Run("Test Update Status Failed When Status is Owned", func(t *testing.T) {
		negotiationBusiness := NewNegotiationBusiness(mockNegotiationDataFailed{})
		row, err := negotiationBusiness.UpdateStatus(1, "owned")
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})
	t.Run("Test Update Status Failed When Status is Cancel", func(t *testing.T) {
		negotiationBusiness := NewNegotiationBusiness(mockNegotiationDataFailed{})
		row, err := negotiationBusiness.UpdateStatus(1, "cancel")
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})
}

func TestDeleteNegotiation(t *testing.T) {
	t.Run("Test Delete Negotiation Success", func(t *testing.T) {
		negotiationBusiness := NewNegotiationBusiness(mockNegotiationData{})
		row, err := negotiationBusiness.DeleteNegotiation(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})
	t.Run("Test Delete Negotiation Failed", func(t *testing.T) {
		negotiationBusiness := NewNegotiationBusiness(mockNegotiationDataFailed{})
		row, err := negotiationBusiness.DeleteNegotiation(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})
}
