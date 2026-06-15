package routes

import (
	"telerad-core-module/internals/controllers"
)

func AuthRoutes() {
	auth := "/auth"

	// Đăng nhập nhân viên (telerad)
	staffCollection := getRoute(v1, staff).Group(auth)
	staffCollection.Post("/token", controllers.StaffLogin)

	// Đăng nhập đối tác telerad (credential phía telerad)
	partnerCollection := getRoute(v1, partner).Group(auth)
	partnerCollection.Post("/token", controllers.TeleradPartnerLogin)
}
