package app

import (
	"fmt"
	"github.com/bitokss/bitok-user-service/constants"
	"github.com/bitokss/bitok-user-service/controllers/v1"
	"github.com/bitokss/bitok-user-service/middlewares/v1"
)

func urlMapper() {
	// users
	e.POST(fmt.Sprintf(constants.V1Prefix, "users"), controllers.UserController.Create, middlewares.OnlyWithPermissions)
	e.GET(fmt.Sprintf(constants.V1Prefix, "users"), controllers.UserController.FindAll)
	e.DELETE(fmt.Sprintf(constants.V1Prefix, "users/:id"), controllers.UserController.Delete)
	e.PUT(fmt.Sprintf(constants.V1Prefix, "users/:id"), controllers.UserController.Update)
	e.GET(fmt.Sprintf(constants.V1Prefix, "users/:id"), controllers.UserController.Find)
	e.POST(fmt.Sprintf(constants.V1Prefix, "users/register"), controllers.UserController.Register)
	e.POST(fmt.Sprintf(constants.V1Prefix, "users/login"), controllers.UserController.Login)
	e.GET(fmt.Sprintf(constants.V1Prefix, "users/byToken/:token"), controllers.UserController.FindByToken)
	e.GET(fmt.Sprintf(constants.V1Prefix, "users/byToken/:username"), controllers.UserController.FindByUsername)
	e.POST(fmt.Sprintf(constants.V1Prefix, "users/resetPassword"), controllers.UserController.ResetPassword)
	e.POST(fmt.Sprintf(constants.V1Prefix, "users/tickRequest"), controllers.UserController.TickRequest)
	// profile
	e.PUT(fmt.Sprintf(constants.V1Prefix, "profile/:username"), controllers.ProfileController.CreateOrUpdate)
	e.GET(fmt.Sprintf(constants.V1Prefix, "profile/:username"), controllers.ProfileController.Find)
	// codes
	e.POST(fmt.Sprintf(constants.V1Prefix, "codes/send"), controllers.CodesController.Send)
	e.POST(fmt.Sprintf(constants.V1Prefix, "codes/verify"), controllers.CodesController.Verify)
	// levels
	e.POST(fmt.Sprintf(constants.V1Prefix, "levels"), controllers.LevelsController.Create)
	e.GET(fmt.Sprintf(constants.V1Prefix, "levels"), controllers.LevelsController.FindAll)
	e.DELETE(fmt.Sprintf(constants.V1Prefix, "levels/:id"), controllers.LevelsController.Delete)
	e.PUT(fmt.Sprintf(constants.V1Prefix, "levels/:id"), controllers.LevelsController.Update)
	e.GET(fmt.Sprintf(constants.V1Prefix, "levels/:id"), controllers.LevelsController.Find)
	// roles
	e.POST(fmt.Sprintf(constants.V1Prefix, "roles"), controllers.RolesController.Create)
	e.GET(fmt.Sprintf(constants.V1Prefix, "roles"), controllers.RolesController.FindAll)
	e.DELETE(fmt.Sprintf(constants.V1Prefix, "roles/:id"), controllers.RolesController.Delete)
	e.PUT(fmt.Sprintf(constants.V1Prefix, "roles/:id"), controllers.RolesController.Update)
	e.GET(fmt.Sprintf(constants.V1Prefix, "roles/:id"), controllers.RolesController.Find)
	// permissions
	e.POST(fmt.Sprintf(constants.V1Prefix, "permissions"), controllers.PermissionsController.Create)
	e.GET(fmt.Sprintf(constants.V1Prefix, "permissions"), controllers.PermissionsController.FindAll)
	e.DELETE(fmt.Sprintf(constants.V1Prefix, "permissions/:id"), controllers.PermissionsController.Delete)
	e.PUT(fmt.Sprintf(constants.V1Prefix, "permissions/:id"), controllers.PermissionsController.Update)
	e.GET(fmt.Sprintf(constants.V1Prefix, "permissions/:id"), controllers.PermissionsController.Find)
}
