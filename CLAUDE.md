# Quy ước codebase

> telerad-core dùng chung kiến trúc phân lớp với his-core (skeleton clone từ his-core, module `telerad-core-module`). Bộ quy ước dưới đây áp dụng khi sinh code mới.

## Mục lục

**Phân lớp & dữ liệu**
1. Tên bảng (`TableXxx`) chỉ xuất hiện ở `base-services` và `repositories`
2. Tên cột (column name) chỉ xuất hiện ở lớp `repositories`
3. Service không tự dựng query — chỉ mở transaction và truyền `tx` xuống
4. Luồng phụ thuộc một chiều: controller → service → (base-service | repository)
5. Hàm ghi ở base-service/repository nhận `tx` từ caller — không tự mở transaction
6. Đặt select validate NGOÀI transaction; chỉ mở transaction khi cần ghi
7. Select quan hệ từ 4 tầng trở lên → dùng GORM (không dùng bun)

**Entity & nghiệp vụ**
8. Init entity phải viết thành hàm `InitNew...` trong base-service
9. Audit field set qua init/update helper — không set tay rải rác
10. Validate cú pháp ở `requests`/controller; service chỉ validate nghiệp vụ
20. Giá trị cố định của cột (enum) định nghĩa DUY NHẤT ở `field-values`

**Response & mapper**
11. Response luôn trả qua object-mappers
12. Mapper là pure function — không `ctx`/`tx`, không query DB
13. Field derived / không thuộc bảng đặt dưới dấu `// separator`
14. Field resolve từ `fieldValues` (lookup có thể nil) phải khai báo `*string`

**Đặt tên & cấu trúc file**
15. Một file = một domain; đặt tên file theo lớp
16. Hàm đặt prefix theo vai trò (`To/InitNew/FindOne/Update...`)
17. Model phục vụ SELECT từ DB khai báo ở `database-query_models`; DTO khác ở `other_models`
18. Model phục vụ FILTER (tham số lọc query) khai báo ở `filter_models`
19. Đọc query/path param phải qua hàm common trong `utils/request_util.go`

---

## Tên bảng (`TableXxx`) chỉ xuất hiện ở `base-services` và `repositories`

**Tên bảng (`repositories.TableXxx`) KHÔNG được phép xuất hiện ở lớp `internals/services/`.** Chỉ được dùng ở `internals/base-services/` và `internals/repositories/`.

### Vì sao
Service là tầng nghiệp vụ — không nên biết bảng nào lưu cái gì. Nếu service trực tiếp gọi `repositories.FindOneByUuidWithSchema[Entity](ctx, tx, schema, repositories.TableXxx, id)` thì:
- Service bị rò rỉ chi tiết tầng dữ liệu (tên bảng).
- Đổi tên bảng / split bảng → phải sửa cả service.
- Khó test mock vì service phụ thuộc constant của repo.

### Cách làm đúng
Bọc lookup vào `baseServices`:

```go
// ✘ Sai — service biết tên bảng
func StaffXxx(...) {
    doi, _ := repositories.FindOneByUuidWithSchema[Entity](
        ctx, tx, schema, repositories.TableXxx, doiUuid)
}

// ✓ Đúng — service chỉ gọi base-service đã typed
func StaffXxx(...) {
    doi, _ := baseServices.FindOneXxxByUuid(ctx, tx, schema, doiUuid)
}

// base-service mới biết tên bảng:
func FindOneXxxByUuid(ctx, tx, schema, id) (*Entity, error) {
    return repositories.FindOneByUuidWithSchema[Entity](
        ctx, tx, schema, repositories.TableXxx, id)
}
```

## Tên cột (column name) chỉ xuất hiện ở lớp `repositories`

**Chuỗi raw column name (`"status = ?"`, `"equipment_uuid = ?"`, ...) trong `Set(...)`, `Where(...)`, `ColumnExpr(...)`, `OrderExpr(...)` KHÔNG được phép xuất hiện ở `internals/services/` hay `internals/base-services/`.** Chỉ `internals/repositories/` mới được phép.

### Vì sao
Tên cột là chi tiết của schema vật lý. Nếu lan ra base-service hoặc service:
- Đổi tên cột → phải sửa nhiều lớp.
- Khó test mock.
- Dễ dẫn tới partial UPDATE bị viết lặp, mỗi nơi 1 kiểu, hard-code khác nhau.

### Quy tắc tổng

| Lớp | Tên bảng (`TableXxx`) | Tên cột (`"col = ?"`) |
|---|---|---|
| `repositories` | ✓ | ✓ |
| `base-services` | ✓ (truyền cho repo helper generic) | ✗ |
| `services` | ✗ | ✗ |

## Response luôn trả qua object-mappers

**Mọi response trả về client PHẢI được dựng qua một hàm mapper nằm trong `internals/object-mappers/`.** Controller / service KHÔNG được tự khởi tạo struct response inline rải rác trong handler.

### Vì sao
- Logic ánh xạ entity → response (derive trạng thái, format, resolve `fieldValues`, gộp quan hệ) tập trung một chỗ → đổi shape response chỉ sửa 1 file.
- Controller mỏng, chỉ điều phối; không lẫn chi tiết trình bày.
- Tái sử dụng được mapper giữa nhiều endpoint, tránh map lặp mỗi nơi một kiểu.

### Cách làm đúng
- File mapper đặt trong `internals/object-mappers/`, đặt tên `<domain>-controller-responses_mapper.go`, package `objectMappers`.
- Hàm mapper đặt tên prefix `To...` và trả về đúng struct response (hoặc slice của nó): `ToStaffGetAXxxResponse(...)`, `ToStaffGetListXxxResponseSlice(...)`.
- Mọi tính toán derived (status, format ngày, ghép tên từ `fieldValues`) nằm trong mapper, không nằm ở controller.

```go
// ✘ Sai — controller tự dựng response inline
return c.JSON(xxxControllerResponses.StaffGetAXxxResponse{
    Uuid: record.Uuid,
    Name: record.Name,
    // ... map tay nhiều field ở đây
})

// ✓ Đúng — trả qua mapper
return c.JSON(objectMappers.ToStaffGetAXxxResponse(*record))
```

## Init entity phải viết thành hàm `InitNew...` trong base-service

**Mọi khởi tạo entity mới (gán default, sinh uuid, set audit field, map từ request) PHẢI nằm trong một hàm `Init...` đặt ở `internals/base-services/`.** Service / controller KHÔNG được `Entity{ ... }` inline để chuẩn bị bản ghi insert.

### Vì sao
- Default value, audit field, business invariant của entity tập trung một chỗ → thêm field mới chỉ sửa hàm init.
- Service không phải biết entity có những field bắt buộc nào.
- Khởi tạo nhất quán giữa các luồng.

### Cách làm đúng
- Đặt tên `InitNew<Entity>(...)` cho 1 bản ghi, `InitMany<Entity>s(...)` cho nhiều bản ghi.
- Hàm nhận request / tham số nghiệp vụ, trả về entity (hoặc slice entity) đã sẵn sàng để base-service insert.

```go
// ✘ Sai — service dựng entity inline
record := clusterEntities.XxxEntity{ Uuid: uuid.New(), Name: request.Name, IsActive: true }

// ✓ Đúng — service gọi init từ base-service
record := baseServices.InitNewXxx(request, true)

// internals/base-services/xxx_base-service.go
func InitNewXxx(request xxxControllerRequests.StaffCreateXxxRequest, isActive bool) clusterEntities.XxxEntity {
    return clusterEntities.XxxEntity{ ... }
}
```

## Đặt select validate NGOÀI transaction; chỉ mở transaction khi cần ghi dữ liệu

**Ưu tiên chạy các câu lệnh SELECT để kiểm tra / validate dữ liệu BÊN NGOÀI transaction.** Chỉ mở transaction khi bắt đầu cần tác động ghi (INSERT / UPDATE / DELETE), và **đóng transaction ngay khi nghiệp vụ ghi cuối cùng kết thúc**.

### Vì sao
- Transaction giữ càng lâu càng giữ lock / connection lâu → giảm throughput, dễ deadlock.
- Phần đọc-validate thường chiếm phần lớn thời gian xử lý; để trong transaction là phí.
- Tách rạch ròi "đọc để quyết định" và "ghi để thay đổi" giúp logic rõ ràng hơn.

### Cách làm đúng

```go
// ✘ Sai — mở transaction từ đầu, validate cũng nằm trong tx
err := db.RunInTx(ctx, func(tx) error {
    record, _ := baseServices.FindOneXxxByUuid(ctx, tx, schema, id)   // chỉ đọc để validate
    if record == nil { return ErrNotFound }
    if record.Status != "pending" { return ErrInvalidState }          // chỉ validate
    return baseServices.UpdateXxx(ctx, tx, schema, ...)               // mới là ghi
})

// ✓ Đúng — validate ngoài tx, chỉ mở tx khi ghi
record, _ := baseServices.FindOneXxxByUuid(ctx, nil, schema, id)      // đọc ngoài transaction
if record == nil { return ErrNotFound }
if record.Status != "pending" { return ErrInvalidState }             // validate xong xuôi

err := db.RunInTx(ctx, func(tx) error {                               // chỉ mở tx khi bắt đầu ghi
    return baseServices.UpdateXxx(ctx, tx, schema, ...)
})                                                                    // đóng tx ngay khi ghi xong
```

Nếu cần đọc lại bên trong transaction để chống race (vd: re-check tồn tại / lock row trước khi ghi) thì vẫn được — nhưng phần validate thuần đọc nên nằm ngoài.

## Select quan hệ từ 4 tầng trở lên → dùng GORM (không dùng bun)

**Câu lệnh nào cần join / preload quan hệ sâu từ 4 tầng trở lên (tính cả bảng gốc) thì dùng GORM thay vì bun.** "Tầng" = mức quan hệ lồng nhau, KHÔNG phải số bảng tham gia.

> Ví dụ 4 tầng: `A` → `A.Bs` → `B.Cs` → `C.Ds`. Đếm theo độ sâu lồng nhau của quan hệ, không phải tổng số bảng join ngang hàng.

### Vì sao
- bun sinh alias tự động cho quan hệ lồng sâu; ở 4 tầng trở lên alias bị quá dài → lỗi truy vấn (identifier too long / alias collision).
- GORM xử lý preload quan hệ sâu ổn định hơn cho các trường hợp này.

### Cách làm đúng
- ≤ 3 tầng quan hệ: dùng bun như bình thường (mặc định của codebase).
- ≥ 4 tầng quan hệ: viết query bằng GORM (đặt ở `internals/repositories/`, vẫn tuân quy tắc tên bảng / tên cột chỉ ở lớp repository).

## Service không tự dựng query — chỉ mở transaction và truyền `tx` xuống

**Service KHÔNG được dùng query builder (`tx.NewSelect()`, `tx.NewUpdate()`, `tx.NewInsert()`, `tx.NewDelete()`, `.Model()`, `.Where()`, ...).** Service chỉ: mở transaction (qua `transactionBun`), điều phối nghiệp vụ, và gọi xuống base-service / repository để thực thi truy vấn.

> Import `bun` ở service để dùng kiểu `bun.Tx` / helper `transactionBun` là HỢP LỆ — cái bị cấm là *dựng câu query* bằng builder trong service.

### Vì sao
- Mọi truy vấn tập trung ở repository → đổi cách query không lan vào nghiệp vụ.
- Khớp với rule tên cột: builder ở service sẽ kéo theo column literal lọt ra ngoài repository.

```go
// ✘ Sai — service tự dựng query
_, err := tx.NewUpdate().Model(&record).Set("status = ?", "done").Where(...).Exec(ctx)

// ✓ Đúng — service mở tx rồi gọi xuống base-service/repository
_, err = transactionBun(ctx, &schema, func(ctx context.Context, tx bun.Tx) (any, error) {
    return nil, baseServices.UpdateXxxStatus(ctx, tx, schema, record.Uuid, "done")
})
```

## Luồng phụ thuộc một chiều: controller → service → (base-service | repository)

**Controller CHỈ gọi service** — không gọi thẳng `baseServices.` hay `repositories.`.
**Service được phép gọi CẢ base-service và repository.** Base-service gọi repository.

```
controller ──▶ service ──▶ base-service ──▶ repository
                   └──────────────────────▶ repository   (được phép)
```

### Vì sao
- Controller mỏng, không nhảy cấp xuống tầng dữ liệu.
- Service là điểm điều phối nghiệp vụ duy nhất; gọi repository trực tiếp được, miễn tuân rule tên bảng / tên cột (không truyền `TableXxx`, không column literal — dùng repo helper đã đóng gói hoặc đi qua base-service).

## Validate cú pháp ở `requests`/controller; service chỉ validate nghiệp vụ

**Validation cú pháp / ràng buộc field (required, định dạng, min/max, enum hợp lệ) đặt ở lớp `internals/requests/` hoặc controller khi parse request.** Service CHỈ validate *nghiệp vụ*: tồn tại bản ghi, trạng thái hợp lệ để chuyển, quyền, ràng buộc liên bảng.

### Vì sao
- Tách lỗi 400 (sai cú pháp request) khỏi lỗi nghiệp vụ (409/422).
- Service không phải lặp lại kiểm tra field rỗng/format.

## Hàm ghi ở base-service/repository nhận `tx` từ caller — không tự mở transaction

**Hàm base-service / repository có tác động ghi PHẢI nhận `tx` từ tham số, không tự gọi `transactionBun` / mở transaction bên trong.** Việc mở và đóng transaction do **service** quyết định.

### Vì sao
- Tránh nested transaction / transaction ẩn.
- Cho phép gộp nhiều thao tác ghi vào cùng một transaction do service điều phối (khớp rule "mở tx muộn, đóng tx sớm").

## Audit field set qua init/update helper — không set tay rải rác

**`created_at` / `created_by` / `updated_at` / `updated_by` phải được set trong hàm `InitNew...` (khi tạo) và trong repository update helper (khi sửa), KHÔNG set tay rải rác ở service/controller.**

### Vì sao
- Nhất quán toàn hệ thống; thêm/đổi audit field chỉ sửa một chỗ (khớp rule `InitNew...` và rule tên cột).

## Field derived / không thuộc bảng đặt dưới dấu `// separator` trong response struct

**Trong struct response, các field map thẳng từ cột bảng đặt phía trên; các field derived / tính toán / không thuộc bảng đặt phía dưới một dòng `// separator`.**

### Vì sao
- Nhìn struct biết ngay field nào là dữ liệu thô, field nào do mapper dựng thêm.

```go
type StaffGetAXxxResponse struct {
    Uuid   uuid.UUID `json:"uuid"`
    Name   string    `json:"name"`
    Status string    `json:"status"`
    // ----- derived (mapper tính, không thuộc bảng) -----
    StatusLabel  string  `json:"status_label"`
    DisplayName  *string `json:"display_name"`
}
```

## Field resolve từ `fieldValues` (lookup có thể nil) phải khai báo `*string`

**Field nào lấy giá trị qua tra cứu `fieldValues` (hoặc bất kỳ lookup có thể không tìm thấy) PHẢI khai báo con trỏ `*string`, không dùng `string`.**

### Vì sao
- Lookup có thể trả nil → dùng `string` sẽ buộc gán giá trị rỗng giả hoặc panic; `*string` phản ánh đúng "không có giá trị".

## Mapper là pure function — không `ctx`/`tx`, không query DB

**Hàm trong `object-mappers` chỉ nhận entity / dto / tham số đã có sẵn và trả về response.** KHÔNG nhận `ctx`/`tx`, KHÔNG query DB, KHÔNG gọi service/repository.

### Vì sao
- Cần dữ liệu phụ thì service load sẵn rồi truyền vào mapper → mapper thuần, dễ test, không gây query ẩn (N+1) lúc dựng response.

## Một file = một domain; đặt tên file theo lớp

**Mỗi domain một file ở mỗi lớp, đặt tên nhất quán:**

| Lớp | Quy ước tên file |
|---|---|
| `services` | `<domain>_service.go` |
| `base-services` | `<domain>_base-service.go` |
| `repositories` | `<domain>_repository.go` |
| `controllers` | `<domain>_controller.go` |
| `object-mappers` | `<domain>-controller-responses_mapper.go` |

### Vì sao
- Code-gen và người đọc đoán được file đích từ domain; tránh gom nhiều domain vào một file.

## Hàm đặt prefix theo vai trò

**Đặt tên hàm theo prefix cố định để đoán được hành vi:**

| Vai trò | Prefix | Ví dụ |
|---|---|---|
| Mapper (entity → response) | `To...` | `ToStaffGetAXxxResponse` |
| Init entity | `InitNew...` / `InitMany...` | `InitNewDepartment` |
| Đọc | `FindOne...` / `FindList...` / `Count...` | `FindOneXxxByUuid` |
| Ghi | `Insert...` / `Update...` / `Delete...` | `UpdateXxxStatus` |

### Vì sao
- Tên hàm tự mô tả tác dụng (đọc hay ghi, một hay nhiều), giảm phải đọc thân hàm.

## Model phục vụ SELECT từ DB khai báo ở `database-query_models`; DTO khác ở `other_models`

**Mọi struct dùng làm đích `Scan`/projection cho một câu SELECT (row model, kết quả join/aggregate không map về entity) PHẢI khai báo ở `internals/models/database-query_models/` (package `databaseQueryModels`).** Các DTO/model trung gian KHÔNG phải kết quả select trực tiếp (kết quả tính toán, dữ liệu ghép tay, value type chia sẻ giữa các tầng) đặt ở `internals/models/other_models/` (package `otherModels`).

### Vì sao
- Tách bạch "hình dạng dữ liệu đọc thẳng từ DB" với "dữ liệu do code tự dựng" → nhìn package biết ngay nguồn gốc.
- Row model nằm cùng một chỗ, dễ rà khi đổi câu query / schema.
- Tránh nhét model query vào `entities` (entity = ánh xạ bảng) hay vào `responses` (response = hình dạng API).

### Cách phân loại

| Loại model | Nơi khai báo |
|---|---|
| Đích Scan của một SELECT (row/projection, join, aggregate) | `database-query_models` (`databaseQueryModels`) |
| Tham số lọc/scope của một query (input WHERE) | `filter_models` (`filterModels`) |
| DTO/model tính toán, không map thẳng từ select | `other_models` (`otherModels`) |
| Ánh xạ 1-1 với bảng | `entities` |
| Hình dạng trả ra API | `responses` |

> Đặt tên file `*_model.go`; struct row-model nên có hậu tố `...RowModel` / `...Model`.

## Model phục vụ filter khai báo ở `filter_models`

**Mọi struct gom tham số lọc/scope cho một query (input của WHERE: khoảng ngày, từ khoá, trạng thái, phạm vi quyền...) PHẢI khai báo ở `internals/models/filter_models/` (package `filterModels`), KHÔNG khai báo trong `repositories/` hay `services/`.**

### Vì sao
- Repository và service đều tham chiếu cùng một kiểu filter mà không bên nào "sở hữu" nó (tránh service phụ thuộc type khai báo trong repo).
- Tham số lọc tập trung một chỗ, dễ thêm điều kiện mới.
- Tách rõ "input lọc" (`filter_models`) với "row kết quả" (`database-query_models`).

```go
// ✘ Sai — struct filter khai báo trong repository
// internals/repositories/telerad-reading-order_repository.go
type ReadingOrderListFilter struct { ... }

// ✓ Đúng — khai báo ở filter_models, repo + service cùng dùng
// internals/models/filter_models/reading-order-list-filter_model.go
package filterModels
type ReadingOrderListFilter struct { ... }
```

## Đọc query/path param phải qua hàm common trong `utils/request_util.go`

**Mọi giá trị đọc từ request param (query string + path param) PHẢI dùng hàm common trong `utils/request_util.go`, KHÔNG gọi thẳng `c.Query(...)` / `c.Params(...)` / `c.QueryArgs()` trong controller.**

### Vì sao
- Parse + default + cờ `allowNull` + xử lý lỗi gom một chỗ → mọi endpoint hành xử nhất quán (cùng kiểu trả `*T`, cùng cách báo lỗi).
- Thêm/đổi cách parse một kiểu (vd: định dạng ngày) chỉ sửa một file.
- Controller mỏng, không lặp `strconv.ParseInt`/`uuid.Parse` mỗi nơi.

### Hàm có sẵn (thêm hàm mới vào đúng file này nếu thiếu kiểu)

| Mục đích | Hàm |
|---|---|
| Phân trang | `GetPaginationParams` |
| Query param (nullable, có lỗi) | `GetInt64FromRequestParam`, `GetInt16FromRequestParam`, `GetIntFromRequestParam`, `GetBoolFromRequestParam`, `GetUuidFromRequestParam`, `GetTimeFromRequestParam` |
| Query param string | `GetStringFromRequestParam` |
| Query param dạng mảng | `GetInt64SliceFromRequestParam`, `GetStringSliceFromRequestParam` |
| Path param | `GetUuidFromRequestPath` |

```go
// ✘ Sai — controller tự đọc param thô
status := c.Query("status", "")
id, _ := uuid.Parse(c.Params("uuid"))

// ✓ Đúng — qua hàm common
status := utils.GetStringFromRequestParam(c, "status")
id, err := utils.GetUuidFromRequestPath(c, "uuid")
```

> Lưu ý: rule này nói về **request param** (query/path). Body request vẫn parse qua `RequestBodyParser` như hiện hành.

## Giá trị cố định của cột (enum) định nghĩa DUY NHẤT ở `field-values`

**Mọi giá trị cố định/enum của một cột (status, type, kind, level, gender, priority...) PHẢI được định nghĩa ở `internals/entities/field-values/` (dưới dạng `ColumnValueString` / `ColumnValueInt16` gắn `Code`), và CHỈ định nghĩa ở đây.** KHÔNG hard-code chuỗi/số raw (`"ACTIVE"`, `"READING"`, `1`...) hay khai báo `const`/`var` giá trị cột rải rác ở `services`/`base-services`/`repositories`/`entities`/`object-mappers`.

### Vì sao
- Một nguồn sự thật duy nhất cho tập giá trị hợp lệ + nhãn hiển thị (`Name`/`ShortName`) → đổi/thêm giá trị chỉ sửa một chỗ.
- Tránh lệch giá trị giữa các nơi (mỗi chỗ gõ tay một kiểu, sai chính tả khó phát hiện).
- Đổ combobox, resolve nhãn, validate "thuộc tập hợp lệ" đều dùng chung định nghĩa + helper (`FromValueAndCodeString`, `GetAllStringTypeByCode`...).

### Cách làm đúng
- Mỗi nhóm giá trị có 1 `Code` (hằng `XXX = "..."`) + các phần tử `ColumnValueString{Value, Code, Name, ShortName}`.
- Nơi khác **tham chiếu** qua `fieldValues.XXX_YYY.Value`, không gõ lại literal.
- Validate giá trị nhập vào bằng `fieldValues.FromValueAndCodeString(value, fieldValues.CODE)` (nil = không hợp lệ).

```go
// ✘ Sai — hard-code / khai báo giá trị cột ngoài field-values
if status == "READING" { ... }
const readingStatusReading = "READING"
record.Status = "DONE"

// ✓ Đúng — định nghĩa ở field-values, nơi khác tham chiếu
// internals/entities/field-values/field-values.go
TELERAD_READING_ORDER_STATUS_READING = ColumnValueString{Value: "READING", Code: TELERAD_READING_ORDER_STATUS, Name: "Đang đọc", ShortName: "Đang đọc"}

// nơi dùng
if status == fieldValues.TELERAD_READING_ORDER_STATUS_READING.Value { ... }
```

> Lưu ý: rule này áp cho **giá trị enum của cột** (tập giá trị cố định, hữu hạn). Không áp cho hằng số kỹ thuật không phải giá trị cột (vd: format layout, key cache, tên header HTTP).
