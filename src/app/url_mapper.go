package app

import (
	"fmt"

	"github.com/bitokss/bitok-user-service/constants"
	"github.com/bitokss/bitok-user-service/controllers/v1"
	"github.com/bitokss/bitok-user-service/middlewares/v1"
)

func urlMapper() {
	// users
	e.POST(fmt.Sprintf(constants.V1Prefix, "users"), controllers.UsersController.Create, middlewares.OnlyWithPermissions([]string{constants.AddUsersPermission}))
	e.GET(fmt.Sprintf(constants.V1Prefix, "users"), controllers.UsersController.FindAll, middlewares.OnlyWithPermissions([]string{constants.FindAllUsersPermission}))
	e.DELETE(fmt.Sprintf(constants.V1Prefix, "users/:id"), controllers.UsersController.Delete, middlewares.OnlyWithPermissions([]string{constants.DeleteUsersPermission}))
	e.PUT(fmt.Sprintf(constants.V1Prefix, "users/:id"), controllers.UsersController.Update, middlewares.OnlyWithPermissions([]string{constants.UpdateUsersPermission}))
	e.GET(fmt.Sprintf(constants.V1Prefix, "users/:id"), controllers.UsersController.Find, middlewares.OnlyWithPermissions([]string{constants.FindUsersPermission}))

	e.POST(fmt.Sprintf(constants.V1Prefix, "users/register"), controllers.UsersController.Register)
	e.POST(fmt.Sprintf(constants.V1Prefix, "users/login"), controllers.UsersController.Login)
	e.GET(fmt.Sprintf(constants.V1Prefix, "users/byToken/:token"), controllers.UsersController.FindByToken)
	e.GET(fmt.Sprintf(constants.V1Prefix, "users/byUsername/:username"), controllers.UsersController.FindByUsername, middlewares.OnlyWithPermissions([]string{constants.FindUsersPermission}))
	e.POST(fmt.Sprintf(constants.V1Prefix, "users/resetPassword"), controllers.UsersController.ResetPassword)
	// profile
	e.PUT(fmt.Sprintf(constants.V1Prefix, "profile/:username"), controllers.ProfileController.CreateOrUpdate, middlewares.OnlyLogin)
	e.GET(fmt.Sprintf(constants.V1Prefix, "profile/:username"), controllers.ProfileController.Find)
	// codes
	e.POST(fmt.Sprintf(constants.V1Prefix, "codes/send"), controllers.CodesController.Send)
	e.POST(fmt.Sprintf(constants.V1Prefix, "codes/verify"), controllers.CodesController.Verify)
	// levels
	e.POST(fmt.Sprintf(constants.V1Prefix, "levels"), controllers.LevelsController.Create, middlewares.OnlyWithPermissions([]string{constants.AddLevelsPermission}))
	e.GET(fmt.Sprintf(constants.V1Prefix, "levels"), controllers.LevelsController.FindAll, middlewares.OnlyWithPermissions([]string{constants.FindAllLevelsPermission}))
	e.DELETE(fmt.Sprintf(constants.V1Prefix, "levels/:id"), controllers.LevelsController.Delete, middlewares.OnlyWithPermissions([]string{constants.DeleteLevelsPermission}))
	e.PUT(fmt.Sprintf(constants.V1Prefix, "levels/:id"), controllers.LevelsController.Update, middlewares.OnlyWithPermissions([]string{constants.UpdateLevelsPermission}))
	e.GET(fmt.Sprintf(constants.V1Prefix, "levels/:id"), controllers.LevelsController.Find, middlewares.OnlyWithPermissions([]string{constants.FindLevelsPermission}))
	// roles
	e.POST(fmt.Sprintf(constants.V1Prefix, "roles"), controllers.RolesController.Create, middlewares.OnlyWithPermissions([]string{constants.AddRolesPermission}))
	e.GET(fmt.Sprintf(constants.V1Prefix, "roles"), controllers.RolesController.FindAll, middlewares.OnlyWithPermissions([]string{constants.FindAllRolesPermission}))
	e.DELETE(fmt.Sprintf(constants.V1Prefix, "roles/:id"), controllers.RolesController.Delete, middlewares.OnlyWithPermissions([]string{constants.DeleteRolesPermission}))
	e.PUT(fmt.Sprintf(constants.V1Prefix, "roles/:id"), controllers.RolesController.Update, middlewares.OnlyWithPermissions([]string{constants.UpdateRolesPermission}))
	e.GET(fmt.Sprintf(constants.V1Prefix, "roles/:id"), controllers.RolesController.Find, middlewares.OnlyWithPermissions([]string{constants.FindRolesPermission}))
	// permissions
	e.POST(fmt.Sprintf(constants.V1Prefix, "permissions"), controllers.PermissionsController.Create, middlewares.OnlyWithPermissions([]string{constants.AddPermissionsPermission}))
	e.GET(fmt.Sprintf(constants.V1Prefix, "permissions"), controllers.PermissionsController.FindAll, middlewares.OnlyWithPermissions([]string{constants.FindAllPermissionsPermission}))
	e.DELETE(fmt.Sprintf(constants.V1Prefix, "permissions/:id"), controllers.PermissionsController.Delete, middlewares.OnlyWithPermissions([]string{constants.DeletePermissionsPermission}))
	e.PUT(fmt.Sprintf(constants.V1Prefix, "permissions/:id"), controllers.PermissionsController.Update, middlewares.OnlyWithPermissions([]string{constants.UpdatePermissionsPermission}))
	e.GET(fmt.Sprintf(constants.V1Prefix, "permissions/:id"), controllers.PermissionsController.Find, middlewares.OnlyWithPermissions([]string{constants.FindPermissionsPermission}))
}
