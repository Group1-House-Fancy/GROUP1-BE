package presentation

import (
	"capstoneproject/features/contractors"
	_requestContractor "capstoneproject/features/contractors/presentation/request"
	"capstoneproject/helper"
	"capstoneproject/middlewares"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ContractorHandler struct {
	contractorBusiness contractors.Business
}

func NewContractorHandler(business contractors.Business) *ContractorHandler {
	return &ContractorHandler{
		contractorBusiness: business,
	}
}

func (h *ContractorHandler) JoinContractor(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
	}
	contractorName := c.FormValue("contractor_name")
	numberSIUJK := c.FormValue("number_siujk")
	phoneNumber := c.FormValue("phone_number")
	email := c.FormValue("email")
	address := c.FormValue("address")
	description := c.FormValue("description")

	urlImage, report, err := helper.AddImageContractor(c)
	if err != nil {
		return c.JSON(report["code"].(int), helper.ResponseFailed(fmt.Sprintf("%s", report["message"])))
	}

	urlCertificate, report, err := helper.AddImageCertificate(c)
	if err != nil {
		return c.JSON(report["code"].(int), helper.ResponseFailed(fmt.Sprintf("%s", report["message"])))
	}
	var newContractor = _requestContractor.Contractor{
		ContractorName:      contractorName,
		NumberSIUJK:         numberSIUJK,
		PhoneNumber:         phoneNumber,
		Email:               email,
		Address:             address,
		Description:         description,
		ImageURL:            urlImage,
		CertificateSIUJKURL: urlCertificate,
		UserID:              uint(idToken),
	}
	dataContractor := _requestContractor.ToCore(newContractor)
	row, err := h.contractorBusiness.CreateContractor(dataContractor)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error()))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to create contractor"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccesNoData("success to create contractor"))
}
