package data

import (
	_bcrypt "capstoneproject/bcrypt"
	"capstoneproject/features/users"
	"fmt"

	"capstoneproject/middlewares"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(conn *gorm.DB) users.Data {
	return &mysqlUserRepository{
		db: conn,
	}
}

func (repo *mysqlUserRepository) PostUser(input users.Core) (row int, err error) {
	passHash, _ := _bcrypt.HashPassword(input.Password)
	user := User{
		FullName:     input.FullName,
		Password:     passHash,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		Address:      input.Address,
		ImageURL:     input.ImageURL,
		IsContractor: false,
	}
	tx := repo.db.Create(&user)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to insert data")
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlUserRepository) AuthUser(email string, password string) (fullName string, token string, imageURL string, isContractor bool, e error) {
	userData := User{}
	repo.db.Where("email = ?", email).First(&userData)
	bool1 := _bcrypt.CheckPasswordHash(password, userData.Password)

	if !bool1 {
		return "", "", "", false, fmt.Errorf("error")
	}
	token, errToken := middlewares.CreateToken(int(userData.ID))
	if errToken != nil {
		return "", "", "", false, errToken
	}
	return token, userData.FullName, userData.ImageURL, userData.IsContractor, nil
}

func (repo *mysqlUserRepository) PutDataUser(id int, data users.Core) (int, error) {
	passwordHash, _ := _bcrypt.HashPassword(data.Password)
	data.Password = passwordHash
	var update = fromCore(data)
	result := repo.db.Model(&User{}).Where("id = ?", id).Updates(&update)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, fmt.Errorf("failed to update data")
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlUserRepository) PutDataUser1(id int, data bool) (int, error) {
	contractorUser := User{
		IsContractor: true,
	}
	result := repo.db.Model(&User{}).Where("id = ?", id).Updates(&contractorUser)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, fmt.Errorf("failed to update data")
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlUserRepository) GetUser(id int) (data users.Core, err error) {
	var userData User

	result := repo.db.First(&userData, id)
	if result.Error != nil {
		return users.Core{}, result.Error
	}
	return userData.toCore(), nil
}

func (repo *mysqlUserRepository) DeleteUser(data int) (row int, err error) {
	var dataUsers User
	result := repo.db.Delete(&dataUsers, data)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to delete user")
	}
	return int(result.RowsAffected), nil
}
