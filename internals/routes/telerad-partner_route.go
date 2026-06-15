package routes

import (
	"telerad-core-module/internals/controllers"
	"telerad-core-module/internals/secure"
)

func TeleradPartnerRoutes() {
	teleradPartner := "/telerad-partner"
	staffCollection := getRoute(v1, staff).Group(teleradPartner, secure.CheckAuthorization())

	staffCollection.Get("", controllers.StaffGetPaginatedTeleradPartners)
	staffCollection.Post("", controllers.StaffCreateTeleradPartner)
	// liệt kê tất cả đối tác (kèm trạng thái) để chọn khi phân quyền đọc phim
	staffCollection.Get("/actions/get-all", controllers.StaffGetAllTeleradPartners)
	staffCollection.Get("/:objectId", controllers.StaffGetATeleradPartner)
	staffCollection.Put("/:objectId", controllers.StaffUpdateTeleradPartner)
	staffCollection.Patch("/:objectId/activate", controllers.StaffActivateTeleradPartner)
	staffCollection.Patch("/:objectId/deactivate", controllers.StaffDeactivateTeleradPartner)

	// cấu hình tài khoản phía partner
	staffCollection.Get("/:objectId/partner-config", controllers.StaffGetTeleradPartnerPartnerConfig)
	staffCollection.Put("/:objectId/partner-config", controllers.StaffUpdateTeleradPartnerPartnerConfig)

	// đổi mật khẩu phía telerad (không cho đổi username)
	staffCollection.Patch("/:objectId/change-password", controllers.StaffChangeTeleradPartnerPassword)

	// cây bên trái: đối tác nhóm theo loại chụp (theo quyền user)
	staffCollection.Get("/actions/get-partners-for-reading", controllers.StaffGetTeleradPartnersForReading)
}
