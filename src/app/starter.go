package app

import (
	"github.com/alidevjimmy/go-rest-utils/crypto"
	"github.com/bitokss/bitok-user-service/constants"
	"github.com/bitokss/bitok-user-service/domains/v1"
	"github.com/bitokss/bitok-user-service/repositories/postgres/v1"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strconv"
)

var (
	e *echo.Echo
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constants.InvalidInputErr)
	}
	return nil
}

func StartApp(port string) {
	e = echo.New()
	// validate inputs using go-playground package
	e.Validator = &Validator{validator: validator.New()}
	urlMapper()
	// initialize postgres and get db instance
	db := repositories.PostgresInit()
	// autoMigrate will automatically create tables using domains
	err := db.AutoMigrate(&domains.Permission{}, &domains.Role{}, &domains.Level{}, &domains.User{}, &domains.Code{})
	if err != nil {
		e.Logger.Error(err)
	}
	// add god to db if not exists :) (main admin)
	addGod(db)
	// start echo server
	e.Logger.Error(e.Start(port))
}




func addGod(db *gorm.DB) {
	phone, password, username, email, firstname, lastname := os.Getenv("GOD_PHONE"),
		os.Getenv("GOD_PASSWORD"),
		os.Getenv("GOD_USERNAME"),
		os.Getenv("GOD_EMAIL"),
		os.Getenv("GOD_FIRSTNAME"),
		os.Getenv("GOD_LASTNAME")
	personnelNum , err := strconv.Atoi(os.Getenv("GOD_PERSONNELNUM"))
	if err != nil {
		e.Logger.Error("GOD_PERSONNELNUM is not valid")
	}
	level := domains.Level{
		Title: "کاربر",
		Color: "#ffffff",
	}
	role := domains.Role{
		Title: "genesis",
		Permissions: constants.Permissions,
	}
	user := domains.User{
		Phone: phone,
		Password: crypto.GenerateSha256(password),
		PersonnelNum: personnelNum,
		Username: username,
		Email: email,
		FirstName: firstname,
		LastName: lastname,
		Roles: []domains.Role{role},
		Level: level,
	}
	if err := db.Where("phone = ?" , user.Phone).FirstOrCreate(&user).Error; err != nil {
		e.Logger.Error(err)
	}
}
