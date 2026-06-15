package routes

import (
	"telerad-core-module/internals/controllers"
	"telerad-core-module/internals/secure"
)

func TeleradReadingOrderRoutes() {
	readingOrder := "/reading-order"

	// đối tác (đã đăng nhập telerad) đẩy ca đọc sang
	partnerCollection := getRoute(v1, partner).Group(readingOrder, secure.CheckAuthorization())
	partnerCollection.Post("", controllers.PartnerCreateReadingOrder)

	// nhân viên (bác sĩ đọc) — màn "Đọc ca"
	staffCollection := getRoute(v1, staff).Group(readingOrder, secure.CheckAuthorization())
	staffCollection.Get("", controllers.StaffGetPaginatedReadingOrders)
	// sinh URL mở PACS viewer cho 1 ca đọc (kèm view-token trong URL hash)
	staffCollection.Get("/:objectId/generate-pacs-viewer-url", controllers.StaffGenerateImagingStudyViewerUrl)
}
