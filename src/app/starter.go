package app

import (
	"github.com/bitokss/bitok-user-service/constants"
	"github.com/bitokss/bitok-user-service/domains/v1"
	"github.com/bitokss/bitok-user-service/repositories/postgres/v1"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
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
	// add full user service permissions
	addPermissions(db)
	// add god role
	addRoles(db)
	// add god to db if not exists :) (main admin)
	addGod(db)
	// start echo server
	e.Logger.Error(e.Start(port))
}

func addPermissions(db *gorm.DB) {
	tx := db.Begin()
	tx.SavePoint("begin")
	for _ , v := range constants.Permissions {
		// check permission exists or not
		permission := &domains.Permission{}
		res := tx.Where("symbol = ?" , v.Symbol).First(&permission)
		if res.Error != nil {
			tx.RollbackTo("begin")
			e.Logger.Error("error in getting data from permissions table")
		}
		if permission != nil {
			continue
		}
		// permission not exists so we should add it to database
		res = tx.Create(v)
		if res.Error != nil {
			tx.RollbackTo("begin")
			e.Logger.Error("error in setting data in permissions table")
		}
	}
}

func addRoles(db *gorm.DB) {

}

func addGod(db *gorm.DB) {
	//phone, password, personnelNum, username, email, firstname, lastname := os.Getenv("GOD_PHONE"),
	//	os.Getenv("GOD_PASSWORD"),
	//	os.Getenv("GOD_PERSONNELNUM"),
	//	os.Getenv("GOD_USERNAME"),
	//	os.Getenv("GOD_EMAIL"),
	//	os.Getenv("GOD_FIRSTNAME"),
	//	os.Getenv("GOD_LASTNAME")

}
