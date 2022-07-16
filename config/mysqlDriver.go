package config

import (
	_mContractors "capstoneproject/features/contractors/data"
	_mHouseImages "capstoneproject/features/houseImages/data"
	_mHouses "capstoneproject/features/houses/data"
	_mUsers "capstoneproject/features/users/data"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dbUsername := os.Getenv("DB_Username")
	dbPassword := os.Getenv("DB_Password")
	dbPort := os.Getenv("DB_Port")
	dbHost := os.Getenv("DB_Host")
	dbName := os.Getenv("DB_Name")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	InitMigrate(db)
	return db
}

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&_mUsers.User{})
	db.AutoMigrate(&_mContractors.Contractor{})
	db.AutoMigrate(&_mHouses.House{})
	db.AutoMigrate(&_mHouseImages.HouseImage{})
}
