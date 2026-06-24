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
	// chi tiết 1 ca đọc (tab chi tiết)
	staffCollection.Get("/:objectId", controllers.StaffGetReadingOrderDetail)
	// nhận ca (UNREAD -> READING) / hủy khóa (READING của mình -> UNREAD)
	staffCollection.Post("/:objectId/actions/receive", controllers.StaffReceiveReadingOrder)
	staffCollection.Post("/:objectId/actions/cancel-lock", controllers.StaffCancelReadingOrderLock)
	// lưu kết quả đọc (result_in_html)
	staffCollection.Post("/:objectId/actions/save-result", controllers.StaffSaveReadingOrderResult)
	// kết thúc & duyệt (READING của mình + có kết quả -> APPROVED)
	staffCollection.Post("/:objectId/actions/end-reading-and-approve", controllers.StaffEndReadingAndApprove)
	// sinh URL mở PACS viewer cho 1 ca đọc (kèm view-token trong URL hash)
	staffCollection.Get("/:objectId/generate-pacs-viewer-url", controllers.StaffGenerateImagingStudyViewerUrl)

	// CÔNG KHAI (HIS / bệnh nhân xem phiếu qua link, màn in của staff cũng dùng) — KHÔNG yêu cầu đăng nhập
	publicCollection := getRoute(v1, public).Group(readingOrder)
	publicCollection.Get("/:objectId/result-sheet", controllers.PublicGetReadingOrderResultSheet)
}
