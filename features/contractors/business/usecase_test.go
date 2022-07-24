package business

import (
	"capstoneproject/features/contractors"
	"capstoneproject/features/users"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockContractorData struct{}

func (mock mockContractorData) PostContractor(data contractors.Core) (int, error) {
	return 1, nil
}

func (mock mockContractorData) ContractorExist(id int, userContractor bool) (int, error) {
	return 1, nil
}

func (mock mockContractorData) SelectAllContractor(limit, offset int) ([]contractors.Core, error) {
	return []contractors.Core{
		{
			ID:                  1,
			ContractorName:      "PT. Bintoro Bangun Indonesia",
			NumberSIUJK:         "1-867624-3173-2-03481",
			ImageURL:            "https://storage.googleapis.com/bucket-project-3/image-contractor.png",
			PhoneNumber:         "0824-6743-1158",
			Email:               "marketing@bintorobuild.co.id",
			Address:             "Jl. Kebembem Raya No.6 - Jakarta Selatan",
			Description:         "Menyediakan jasa kontraktor yang membantu Anda dalam menciptakan bangunan yang sesuai dengan impian Anda. Karena impian Anda adalah semangat kami untuk bekerja",
			CertificateSIUJKURL: "https://storage.googleapis.com/bucket-project-3/certificate-contractor.png",
		},
	}, nil
}

func (mock mockContractorData) SelectContractor(idCtr int) (contractors.Core, error) {
	return contractors.Core{
		ID:                  1,
		ContractorName:      "PT. Bintoro Bangun Indonesia",
		NumberSIUJK:         "1-867624-3173-2-03481",
		ImageURL:            "https://storage.googleapis.com/bucket-project-3/image-contractor.png",
		PhoneNumber:         "0824-6743-1158",
		Email:               "marketing@bintorobuild.co.id",
		Address:             "Jl. Kebembem Raya No.6 - Jakarta Selatan",
		Description:         "Menyediakan jasa kontraktor yang membantu Anda dalam menciptakan bangunan yang sesuai dengan impian Anda. Karena impian Anda adalah semangat kami untuk bekerja",
		CertificateSIUJKURL: "https://storage.googleapis.com/bucket-project-3/certificate-contractor.png",
	}, nil
}

func (mock mockContractorData) DeleteContractor(idUser int) (int, error) {
	return 1, nil
}

func (mock mockContractorData) UpdateContractor(idCtr int, idUser int, data contractors.Core) (row int, err error) {
	return 1, nil
}

func (mock mockContractorData) CountContractorData() (count int, err error) {
	return 6, nil
}

func (mock mockContractorData) SelectOwnContractor(idUser int) (contractors.Core, error) {
	return contractors.Core{
		ID:                  1,
		ContractorName:      "PT. Bintoro Bangun Indonesia",
		NumberSIUJK:         "1-867624-3173-2-03481",
		ImageURL:            "https://storage.googleapis.com/bucket-project-3/image-contractor.png",
		PhoneNumber:         "0824-6743-1158",
		Email:               "marketing@bintorobuild.co.id",
		Address:             "Jl. Kebembem Raya No.6 - Jakarta Selatan",
		Description:         "Menyediakan jasa kontraktor yang membantu Anda dalam menciptakan bangunan yang sesuai dengan impian Anda. Karena impian Anda adalah semangat kami untuk bekerja",
		CertificateSIUJKURL: "https://storage.googleapis.com/bucket-project-3/certificate-contractor.png",
	}, nil
}

type mockContractorDataFailed struct{}

func (mock mockContractorDataFailed) PostContractor(data contractors.Core) (int, error) {
	return 0, fmt.Errorf("failed to insert data")
}

func (mock mockContractorDataFailed) ContractorExist(id int, userContractor bool) (int, error) {
	return 0, fmt.Errorf("failed to get data")
}

func (mock mockContractorDataFailed) SelectAllContractor(limit, offset int) ([]contractors.Core, error) {
	return nil, fmt.Errorf("failed to get all data")
}

func (mock mockContractorDataFailed) SelectContractor(idCtr int) (contractors.Core, error) {
	return contractors.Core{}, fmt.Errorf("failed to get data")
}

func (mock mockContractorDataFailed) DeleteContractor(idUser int) (int, error) {
	return 0, fmt.Errorf("failed to delete data")
}

func (mock mockContractorDataFailed) UpdateContractor(idCtr int, idUser int, data contractors.Core) (row int, err error) {
	return 0, fmt.Errorf("failed to update data")
}

func (mock mockContractorDataFailed) CountContractorData() (count int, err error) {
	return 0, fmt.Errorf("failed to count data")
}

func (mock mockContractorDataFailed) SelectOwnContractor(idUser int) (contractors.Core, error) {
	return contractors.Core{
		ID:                  1,
		ContractorName:      "PT. Bintoro Bangun Indonesia",
		NumberSIUJK:         "1-867624-3173-2-03481",
		ImageURL:            "https://storage.googleapis.com/bucket-project-3/image-contractor.png",
		PhoneNumber:         "0824-6743-1158",
		Email:               "marketing@bintorobuild.co.id",
		Address:             "Jl. Kebembem Raya No.6 - Jakarta Selatan",
		Description:         "Menyediakan jasa kontraktor yang membantu Anda dalam menciptakan bangunan yang sesuai dengan impian Anda. Karena impian Anda adalah semangat kami untuk bekerja",
		CertificateSIUJKURL: "https://storage.googleapis.com/bucket-project-3/certificate-contractor.png",
	}, nil
}

type mockUserData struct{}

func (mock mockUserData) PostUser(data users.Core) (int, error) {
	return 1, nil
}

func (mock mockUserData) AuthUser(email string, password string) (string, string, string, bool, error) {
	return "andri gunawan", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTg1MDM2MjQsInVzZXJJZCI6N30.DFxWe18OoFhRAIt41dIDWPQPORI8S5rAkTylJiolnbc", "https://storage.googleapis.com/bucket-project-capstone/default_profile.png", false, nil
}

func (mock mockUserData) PutDataUser(id int, data users.Core) (int, error) {
	return 1, nil
}

func (mock mockUserData) PutDataUser1(id int, data bool) (int, error) {
	return 1, nil
}

func (mock mockUserData) GetUser(id int) (users.Core, error) {
	return users.Core{
		ID: 1, FullName: "Andri G", Email: "andri@gmail.com", PhoneNumber: "081234", Address: "Jln. Urip", ImageURL: "https://storage.com/profile.png", IsContractor: false, CreatedAt: time.Time{},
	}, nil
}

func (mockUserData) DeleteUser(id int) (int, error) {
	return 1, nil
}

type mockUserDataFailed struct{}

func (mock mockUserDataFailed) PostUser(data users.Core) (int, error) {
	return 0, fmt.Errorf("Failed to select data")
}

func (mock mockUserDataFailed) AuthUser(email string, password string) (string, string, string, bool, error) {
	return "", "", "", false, fmt.Errorf("email or password incorrect")
}

func (mock mockUserDataFailed) PutDataUser(id int, data users.Core) (int, error) {
	return 0, fmt.Errorf("failed to update user")
}

func (mock mockUserDataFailed) PutDataUser1(id int, data bool) (int, error) {
	return 0, fmt.Errorf("failed to update user")
}

func (mock mockUserDataFailed) GetUser(id int) (users.Core, error) {
	return users.Core{}, fmt.Errorf("failed to get data")
}

func (mockUserDataFailed) DeleteUser(id int) (int, error) {
	return 0, fmt.Errorf("failed to delete data")
}

func TestCreateContractor(t *testing.T) {
	t.Run("Test Create Contractor Success", func(t *testing.T) {
		contractorBusiness := NewContractorBusiness(mockContractorData{}, mockUserData{})
		var data = contractors.Core{
			ContractorName:      "PT. Bangun Indonesia",
			NumberSIUJK:         "127-066-0",
			ImageURL:            "contractor.png",
			CertificateSIUJKURL: "certificate.png",
			PhoneNumber:         "081234587",
			Email:               "bangunindonesia@gmail.com",
			Address:             "Jalan Raya Bogor-Jakarta, Jakarta",
			Description:         "Membangun Rumah, Gedung, dan Perakntoran dengan material terbaik",
			User: contractors.User{
				ID: 1,
			},
		}
		row, err := contractorBusiness.CreateContractor(data)
		assert.NotNil(t, err)
		assert.Equal(t, -1, row)
	})
	t.Run("Test Create Contractor Failed", func(t *testing.T) {
		contractorBusiness := NewContractorBusiness(mockContractorDataFailed{}, mockUserDataFailed{})
		var data = contractors.Core{
			ContractorName:      "PT. Bangun Indonesia",
			NumberSIUJK:         "127-066-0",
			ImageURL:            "contractor.png",
			CertificateSIUJKURL: "certificate.png",
			PhoneNumber:         "081234587",
			Email:               "bangunindonesia@gmail.com",
			Address:             "Jalan Raya Bogor-Jakarta, Jakarta",
			Description:         "Membangun Rumah, Gedung, dan Perakntoran dengan material terbaik",
			User: contractors.User{
				ID: 1,
			},
		}
		row, err := contractorBusiness.CreateContractor(data)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})
	t.Run("Test Create Contractor Failed Whenn ContractorName is Empty", func(t *testing.T) {
		contractorBusiness := NewContractorBusiness(mockContractorDataFailed{}, mockUserDataFailed{})
		var data = contractors.Core{
			NumberSIUJK:         "127-066-0",
			ImageURL:            "contractor.png",
			CertificateSIUJKURL: "certificate.png",
			PhoneNumber:         "081234587",
			Email:               "bangunindonesia@gmail.com",
			Address:             "Jalan Raya Bogor-Jakarta, Jakarta",
			Description:         "Membangun Rumah, Gedung, dan Perakntoran dengan material terbaik",
			User: contractors.User{
				ID: 1,
			},
		}
		row, err := contractorBusiness.CreateContractor(data)
		assert.NotNil(t, err)
		assert.Equal(t, -1, row)
	})
}

func TestGetAllContractor(t *testing.T) {
	t.Run("Test Get All Contractor Success", func(t *testing.T) {
		contractorBusiness := NewContractorBusiness(mockContractorData{}, mockUserData{})
		data, totalPage, err := contractorBusiness.GetAllContractor(0, 0)
		assert.Nil(t, err)
		assert.Equal(t, []contractors.Core{
			{
				ID:                  1,
				ContractorName:      "PT. Bintoro Bangun Indonesia",
				NumberSIUJK:         "1-867624-3173-2-03481",
				ImageURL:            "https://storage.googleapis.com/bucket-project-3/image-contractor.png",
				PhoneNumber:         "0824-6743-1158",
				Email:               "marketing@bintorobuild.co.id",
				Address:             "Jl. Kebembem Raya No.6 - Jakarta Selatan",
				Description:         "Menyediakan jasa kontraktor yang membantu Anda dalam menciptakan bangunan yang sesuai dengan impian Anda. Karena impian Anda adalah semangat kami untuk bekerja",
				CertificateSIUJKURL: "https://storage.googleapis.com/bucket-project-3/certificate-contractor.png",
			},
		}, data)
		assert.Equal(t, 1, totalPage)
	})
	t.Run("Test Get All Contractor Success When Limit is Odd", func(t *testing.T) {
		contractorBusiness := NewContractorBusiness(mockContractorData{}, mockUserData{})
		data, totalPage, err := contractorBusiness.GetAllContractor(3, 0)
		assert.Nil(t, err)
		assert.Equal(t, []contractors.Core{
			{
				ID:                  1,
				ContractorName:      "PT. Bintoro Bangun Indonesia",
				NumberSIUJK:         "1-867624-3173-2-03481",
				ImageURL:            "https://storage.googleapis.com/bucket-project-3/image-contractor.png",
				PhoneNumber:         "0824-6743-1158",
				Email:               "marketing@bintorobuild.co.id",
				Address:             "Jl. Kebembem Raya No.6 - Jakarta Selatan",
				Description:         "Menyediakan jasa kontraktor yang membantu Anda dalam menciptakan bangunan yang sesuai dengan impian Anda. Karena impian Anda adalah semangat kami untuk bekerja",
				CertificateSIUJKURL: "https://storage.googleapis.com/bucket-project-3/certificate-contractor.png",
			},
		}, data)
		assert.Equal(t, 2, totalPage)
	})
	t.Run("Test Get All Contractor Success When Limit is Even", func(t *testing.T) {
		contractorBusiness := NewContractorBusiness(mockContractorData{}, mockUserData{})
		data, totalPage, err := contractorBusiness.GetAllContractor(2, 0)
		assert.Nil(t, err)
		assert.Equal(t, []contractors.Core{
			{
				ID:                  1,
				ContractorName:      "PT. Bintoro Bangun Indonesia",
				NumberSIUJK:         "1-867624-3173-2-03481",
				ImageURL:            "https://storage.googleapis.com/bucket-project-3/image-contractor.png",
				PhoneNumber:         "0824-6743-1158",
				Email:               "marketing@bintorobuild.co.id",
				Address:             "Jl. Kebembem Raya No.6 - Jakarta Selatan",
				Description:         "Menyediakan jasa kontraktor yang membantu Anda dalam menciptakan bangunan yang sesuai dengan impian Anda. Karena impian Anda adalah semangat kami untuk bekerja",
				CertificateSIUJKURL: "https://storage.googleapis.com/bucket-project-3/certificate-contractor.png",
			},
		}, data)
		assert.Equal(t, 3, totalPage)
	})
	t.Run("Test Get All Contractor Failed", func(t *testing.T) {
		contractorBusiness := NewContractorBusiness(mockContractorDataFailed{}, mockUserDataFailed{})
		data, totalPage, err := contractorBusiness.GetAllContractor(2, 0)
		assert.NotNil(t, err)
		assert.Nil(t, data)
		assert.Equal(t, 0, totalPage)
	})
}

func TestGetContractor(t *testing.T) {
	t.Run("Test Get Contractor Success", func(t *testing.T) {
		contractorBusiness := NewContractorBusiness(mockContractorData{}, mockUserData{})
		data, err := contractorBusiness.GetContractor(1)
		assert.Nil(t, err)
		assert.Equal(t, contractors.Core{
			ID:                  1,
			ContractorName:      "PT. Bintoro Bangun Indonesia",
			NumberSIUJK:         "1-867624-3173-2-03481",
			ImageURL:            "https://storage.googleapis.com/bucket-project-3/image-contractor.png",
			PhoneNumber:         "0824-6743-1158",
			Email:               "marketing@bintorobuild.co.id",
			Address:             "Jl. Kebembem Raya No.6 - Jakarta Selatan",
			Description:         "Menyediakan jasa kontraktor yang membantu Anda dalam menciptakan bangunan yang sesuai dengan impian Anda. Karena impian Anda adalah semangat kami untuk bekerja",
			CertificateSIUJKURL: "https://storage.googleapis.com/bucket-project-3/certificate-contractor.png",
		}, data)
	})
	t.Run("Test Get Contractor Failed", func(t *testing.T) {
		contractorBusiness := NewContractorBusiness(mockContractorDataFailed{}, mockUserDataFailed{})
		data, err := contractorBusiness.GetContractor(1)
		assert.NotNil(t, err)
		assert.Equal(t, contractors.Core{}, data)
	})
}

func TestDeleteContractor(t *testing.T) {
	t.Run("Test Delete Contractor Success", func(t *testing.T) {
		contractorBusiness := NewContractorBusiness(mockContractorData{}, mockUserData{})
		row, err := contractorBusiness.DeleteContractor(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})
	t.Run("Test Delete Contractor Failed", func(t *testing.T) {
		contractorBusiness := NewContractorBusiness(mockContractorDataFailed{}, mockUserDataFailed{})
		row, err := contractorBusiness.DeleteContractor(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})
}

func TestPutContractor(t *testing.T) {
	t.Run("Test Put Contractor Success", func(t *testing.T) {
		contractorBusiness := NewContractorBusiness(mockContractorData{}, mockUserData{})
		var update = contractors.Core{
			ContractorName:      "PT. Bangun Indonesia",
			NumberSIUJK:         "127-066-0",
			ImageURL:            "contractor.png",
			CertificateSIUJKURL: "certificate.png",
			PhoneNumber:         "081234587",
			Email:               "bangunindonesia@gmail.com",
			Address:             "Jalan Raya Bogor-Jakarta, Jakarta",
			Description:         "Membangun Rumah, Gedung, dan Perakntoran dengan material terbaik",
		}
		row, err := contractorBusiness.PutContractor(1, 1, update)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})
	t.Run("Test Put Contractor Failed", func(t *testing.T) {
		contractorBusiness := NewContractorBusiness(mockContractorDataFailed{}, mockUserDataFailed{})
		var update = contractors.Core{
			ContractorName:      "PT. Bangun Indonesia",
			NumberSIUJK:         "127-066-0",
			ImageURL:            "contractor.png",
			CertificateSIUJKURL: "certificate.png",
			PhoneNumber:         "081234587",
			Email:               "bangunindonesia@gmail.com",
			Address:             "Jalan Raya Bogor-Jakarta, Jakarta",
			Description:         "Membangun Rumah, Gedung, dan Perakntoran dengan material terbaik",
		}
		row, err := contractorBusiness.PutContractor(1, 1, update)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})
	t.Run("Test Put Contractor Failed When ContratorName is Empty", func(t *testing.T) {
		contractorBusiness := NewContractorBusiness(mockContractorDataFailed{}, mockUserDataFailed{})
		var update = contractors.Core{
			NumberSIUJK:         "127-066-0",
			ImageURL:            "contractor.png",
			CertificateSIUJKURL: "certificate.png",
			PhoneNumber:         "081234587",
			Email:               "bangunindonesia@gmail.com",
			Address:             "Jalan Raya Bogor-Jakarta, Jakarta",
			Description:         "Membangun Rumah, Gedung, dan Perakntoran dengan material terbaik",
		}
		row, err := contractorBusiness.PutContractor(1, 1, update)
		assert.NotNil(t, err)
		assert.Equal(t, -1, row)
	})
}

func TestGetOwnContractor(t *testing.T) {
	t.Run("Test Get Own Contractor Success", func(t *testing.T) {
		contractorBusiness := NewContractorBusiness(mockContractorData{}, mockUserData{})
		data, err := contractorBusiness.GetOwnContractor(1)
		assert.Nil(t, err)
		assert.Equal(t, contractors.Core{
			ID:                  1,
			ContractorName:      "PT. Bintoro Bangun Indonesia",
			NumberSIUJK:         "1-867624-3173-2-03481",
			ImageURL:            "https://storage.googleapis.com/bucket-project-3/image-contractor.png",
			PhoneNumber:         "0824-6743-1158",
			Email:               "marketing@bintorobuild.co.id",
			Address:             "Jl. Kebembem Raya No.6 - Jakarta Selatan",
			Description:         "Menyediakan jasa kontraktor yang membantu Anda dalam menciptakan bangunan yang sesuai dengan impian Anda. Karena impian Anda adalah semangat kami untuk bekerja",
			CertificateSIUJKURL: "https://storage.googleapis.com/bucket-project-3/certificate-contractor.png",
		}, data)
	})
	t.Run("Test Get Own Contractor Failed", func(t *testing.T) {
		contractorBusiness := NewContractorBusiness(mockContractorDataFailed{}, mockUserDataFailed{})
		data, err := contractorBusiness.GetOwnContractor(1)
		assert.Nil(t, err)
		assert.NotNil(t, contractors.Core{}, data)
	})
}
