package presentation

import (
	"capstoneproject/features/users"
	_requestUser "capstoneproject/features/users/presentation/request"
	_responseUser "capstoneproject/features/users/presentation/response"
	_helper "capstoneproject/helpers"
	"capstoneproject/middlewares"
	"fmt"

	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userBusiness users.Business
}

func NewUserHandler(business users.Business) *UserHandler {
	return &UserHandler{
		userBusiness: business,
	}
}

func (h *UserHandler) AddUser(c echo.Context) error {
	link, report, err := _helper.AddImageUser(c)
	if err != nil {
		return c.JSON(report["code"].(int), _helper.ResponseFailed(fmt.Sprintf("%s", report["message"])))
	}
	var dataUser = _requestUser.User{
		ImageURL: link,
	}
	errBind := c.Bind(&dataUser)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
	}
	v := validator.New()
	errValidator := v.Struct(dataUser)
	errFullName := v.Var(dataUser.FullName, "required,alpha")
	if errFullName != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("fullname can only contains alphabet"))
	}
	if len(dataUser.FullName) == 0 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("fullname must be filled"))
	}
	errEmail := v.Var(dataUser.Email, "required,email")
	if errEmail != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid format email"))
	}
	if len(dataUser.Password) == 0 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("password must be filled"))
	}
	errPhoneNumber := v.Var(dataUser.PhoneNumber, "required,numeric")
	if errPhoneNumber != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("phone number must be in numeric"))
	}
	if errValidator != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed(errValidator.Error()))
	}
	dataUsr := _requestUser.ToCore(dataUser)
	row, err := h.userBusiness.InsertUser(dataUsr)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("email already exist"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, _helper.ResponseSuccesNoData("Succes to insert data"))
}

func (h *UserHandler) Login(c echo.Context) error {
	var userLogin _requestUser.User
	errLog := c.Bind(&userLogin)
	if errLog != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to bind data, check your input"))
	}
	// v := validator.New()
	// errValidator := v.Struct(userLogin)
	// if errValidator != nil {
	// 	return c.JSON(http.StatusBadRequest, _helper.ResponseFailed(errValidator.Error()))
	// }
	token, fullName, imageURL, isContractor, e := h.userBusiness.LoginUser(userLogin.Email, userLogin.Password)
	if e != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("email incorrect"))
	}
	data := map[string]interface{}{
		"full_name":     fullName,
		"token":         token,
		"image_url":     imageURL,
		"is_contractor": isContractor,
	}
	return c.JSON(http.StatusOK, _helper.ResponseSuccesWithData("LOGIN SUCCES", data))
}

func (h *UserHandler) EditData(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}

	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, _helper.ResponseFailed("unauthorized"))
	}
	fullName := c.FormValue("full_name")
	password := c.FormValue("password")
	email := c.FormValue("email")
	phoneNumber := c.FormValue("phone_number")
	address := c.FormValue("address")

	link, report, err := _helper.AddImageUser(c)
	if err != nil {
		return c.JSON(report["code"].(int), _helper.ResponseFailed(fmt.Sprintf("%s", report["message"])))
	}
	data, _ := h.userBusiness.SelectUser(idToken)
	var user = _requestUser.User{
		FullName:    fullName,
		Email:       email,
		Password:    password,
		PhoneNumber: phoneNumber,
		Address:     address,
		ImageURL:    link,
	}

	v := validator.New()
	errValidator := v.Struct(user)

	if user.FullName == "" {
		user.FullName = data.FullName
	}
	errFullName := v.Var(user.FullName, "required,alpha")
	if errFullName != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("fullname can only contains alphabet"))
	}
	if user.Email == "" {
		user.Email = data.Email
	}

	errEmail := v.Var(user.Email, "required,email")
	if errEmail != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid format email"))
	}

	if user.Password == "" {
		user.Password = data.Password
	}

	if user.PhoneNumber == "" {
		user.PhoneNumber = data.PhoneNumber
	}
	errPhoneNumber := v.Var(user.PhoneNumber, "required,numeric")
	if errPhoneNumber != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("phone number must be in numeric"))
	}

	if user.Address == "" {
		user.Address = data.Address
	}

	if errValidator != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed(errValidator.Error()))
	}
	result, err := h.userBusiness.UpdateDataUser(idToken, _requestUser.ToCore(user))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed(err.Error()))
	}
	if result == 0 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to update data"))
	}
	// if result == -1 {
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"message": err.Error(),
	// 	})
	// }
	return c.JSON(http.StatusOK, _helper.ResponseSuccesNoData("success update data"))
}

func (h *UserHandler) GetUser(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}

	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, _helper.ResponseFailed("unauthorized"))
	}
	result, err := h.userBusiness.SelectUser(idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get data user"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseSuccesWithData("success", _responseUser.FromCore(result)))
}

func (h *UserHandler) DeleteDataUser(c echo.Context) error {
	idTok, errDel := middlewares.ExtractToken(c)
	if errDel != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}

	if idTok == 0 {
		return c.JSON(http.StatusUnauthorized, _helper.ResponseFailed("Unauthorized"))
	}
	_, err := h.userBusiness.DeleteUser(idTok)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to delete user"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseSuccesNoData("success to delete user"))
}
