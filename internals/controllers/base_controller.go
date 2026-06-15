package controllers

// base_controller chứa các helper dùng chung cho controller (vd: lấy company/schema
// từ JWT). Skeleton chưa có domain (CompanyEntity, services...) nên để trống; thêm dần
// khi domain xuất hiện. Mẫu tham khảo từ his-core:
//
//	func GetCompanyUuidAndSchemaFromJwt(c *fiber.Ctx) (uuid.UUID, string, error) { ... }
//	func GetCompanyFromJwt(c *fiber.Ctx) (masterEntities.CompanyEntity, error) { ... }
