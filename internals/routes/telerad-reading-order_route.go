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
}
