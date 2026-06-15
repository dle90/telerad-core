package routes

import (
	"telerad-core-module/internals/controllers"
	"telerad-core-module/internals/secure"
)

func StaffAccountRoutes() {
	staffAccount := "/staff-account"
	staffCollection := getRoute(v1, staff).Group(staffAccount, secure.CheckAuthorization())

	// self endpoints - đặt trước /:objectId để không bị nuốt vào param
	staffCollection.Get("/actions/me", controllers.UserGetMe)
	staffCollection.Patch("/actions/change-password", controllers.UserChangePassword)

	staffCollection.Get("", controllers.StaffGetPaginatedStaffAccounts)
	staffCollection.Post("", controllers.StaffCreateStaffAccount)
	staffCollection.Get("/:objectId", controllers.StaffGetAStaffAccount)
	staffCollection.Put("/:objectId", controllers.StaffUpdateStaffAccount)
	staffCollection.Patch("/:objectId/activate", controllers.StaffActivateStaffAccount)
	staffCollection.Patch("/:objectId/deactivate", controllers.StaffDeactivateStaffAccount)

	// phân quyền đọc phim (modalities + đối tác telerad) / phân roles
	staffCollection.Patch("/:objectId/reading-permission", controllers.StaffAssignReadingPermission)
	staffCollection.Patch("/:objectId/roles", controllers.StaffAssignRoles)

	// cấp tài khoản (staff chưa có username) / reset mật khẩu (staff đã có username)
	staffCollection.Post("/:objectId/create-account", controllers.StaffCreateAccount)
	staffCollection.Patch("/:objectId/reset-password", controllers.StaffResetStaffAccountPassword)
}
