package constants

import "github.com/bitokss/bitok-user-service/domains/v1"

const (
	// users
	AddUsersPermission = "add:users"
	FindAllUsersPermission = "findAll:users"
	DeleteUsersPermission = "delete:users"
	UpdateUsersPermission = "update:users"
	FindUsersPermission = "find:users"
	// profile
	CreateOrUpdateProfilePermission = "createOrUpdate:profile"
	// levels
	AddLevelsPermission = "add:levels"
	FindAllLevelsPermission = "findAll:levels"
	DeleteLevelsPermission = "delete:levels"
	UpdateLevelsPermission = "update:levels"
	FindLevelsPermission = "find:levels"
	// roles
	AddRolesPermission = "add:roles"
	FindAllRolesPermission = "findAll:roles"
	DeleteRolesPermission = "delete:roles"
	UpdateRolesPermission = "update:roles"
	FindRolesPermission = "find:roles"
	//permissions
	AddPermissionsPermission = "add:permissions"
	FindAllPermissionsPermission = "findAll:permissions"
	DeletePermissionsPermission = "delete:permissions"
	UpdatePermissionsPermission = "update:permissions"
	FindPermissionsPermission = "find:permissions"
)

var Permissions = []domains.Permission{
	{
		Symbol: AddUsersPermission,
		Title: "افزودن کاربر جدید",
	},
	{
		Symbol: FindAllUsersPermission,
		Title: "مشاهده کابرا",
	},
	{
		Symbol: DeleteUsersPermission,
		Title: "حذف کاربر",
	},
	{
		Symbol: UpdateUsersPermission,
		Title: "ویرایش کاربر",
	},
	{
		Symbol: FindUsersPermission,
		Title: "مشاهده کاربر خاص",
	},
	{
		Symbol: CreateOrUpdateProfilePermission,
		Title: "ویرایش صفحه شخصی",
	},
	{
		Symbol: AddLevelsPermission,
		Title: "افزودن سطح جدید",
	},
	{
		Symbol: FindAllLevelsPermission,
		Title: "مشاهده سطح ها",
	},
	{
		Symbol: DeleteLevelsPermission,
		Title: "حذف سطح",
	},
	{
		Symbol: UpdateLevelsPermission,
		Title: "ویرایش سطح",
	},
	{
		Symbol: FindLevelsPermission,
		Title: "مشاهده سطح خاص",
	},
	{
		Symbol: AddRolesPermission,
		Title: "افزودن نقش جدید",
	},
	{
		Symbol: FindAllRolesPermission,
		Title: "مشاهده نقش ها",
	},
	{
		Symbol: DeleteRolesPermission,
		Title: "حذف نقش",
	},
	{
		Symbol: UpdateRolesPermission,
		Title: "ویرایش نقش",
	},
	{
		Symbol: FindRolesPermission,
		Title: "مشاهده نقش خاص",
	},
	{
		Symbol: AddPermissionsPermission,
		Title: "افزودن دسترسی جدید",
	},
	{
		Symbol: FindAllPermissionsPermission,
		Title: "مشاهده دسترسی ها",
	},
	{
		Symbol: DeletePermissionsPermission,
		Title: "حذف دسترسی",
	},
	{
		Symbol: UpdatePermissionsPermission,
		Title: "ویرایش دسترسی",
	},
	{
		Symbol: FindPermissionsPermission,
		Title:  "مشاهده دسترسی خاص",
	},
}