#### hệ thống quản lý dịch vụ auth
Hệ thống được sinh ra đáp ứng yêu cầu:
1. Đăng ký
2. Đăng nhập
3. ...

### Swagger URL:
# Reference: https://articles.wesionary.team/automatically-generate-restful-api-documentation-in-golang-76927f8f8935
# URL: http://localhost:8080/auth/swagger/index.html

# Hướng dẫn thêm api:
- B1: Thêm bằng cơm vào trước tên hàm. 
	VD1: 
	// CreateNews ... Create News
	// @Summary Create new news based on paramters
	// @Description Create new news
	// @Tags News
	// @Accept json
	// @Param news body models.News true "News Data"
	// @Success 200 {object} object
	// @Failure 400,500 {object} object
	// @Router / [post]
	func CreateNews(c *fiber.Ctx) error {
		...
	}
	VD2: 
	// GetNewsByID ... Get the news by id
	// @Summary Get one news
	// @Description get news by ID
	// @Tags News
	// @Param id path string true "News ID"
	// @Success 200 {object} models.News
	// @Failure 400,404 {object} object
	// @Router /{id} [get]
	func GetANews(c *fiber.Ctx) error {
	...
	}
	VD3: 
	// GetNews ... Get all news
	// @Summary Get all news
	// @Description get all news
	// @Tags News
	// @Success 200 {array} models.News
	// @Failure 404 {object} object
	// @Router / [get]
	func GetAllNews(c *fiber.Ctx) error {
		...
	}
- B2: Chạy câu lệnh bên dưỡi để tự động generate api trên swagger
	$ ~/go/bin/swag init

- B3: Build lại và xem thành quả http://localhost:8080/...




