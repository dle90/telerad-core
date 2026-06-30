package partnerIntegrationService

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"telerad-core-module/internals/entities"
	partnerRequests "telerad-core-module/internals/integration-services/reading-partner/requests"
	partnerResponses "telerad-core-module/internals/integration-services/reading-partner/responses"

	"github.com/BeeTechHub/go-common/logger"
	"github.com/google/uuid"
)

// Đối xứng với his-core teleradIntegrationService (chiều his-core -> telerad). Ở đây
// telerad là bên GỌI, đối tác (his-core) là bên NHẬN. Credential + URL callback lấy
// từ telerad_partner (PartnerUsername/PartnerPassword/CallbackUrl) — mỗi đối tác 1 bộ.

type partnerToken struct {
	AccessToken string
	ExpiresAt   time.Time
}

// Đường dẫn API telerad — GENERIC, KHÔNG gắn với his-core cụ thể, để tích hợp được nhiều
// phần mềm HIS khác nhau. Phần đặc thù của từng HIS (vd his-core luôn có /services/his-core)
// nằm TRONG CallbackUrl cấu hình theo từng đối tác; telerad chỉ nối CallbackUrl + path dưới.
// Vd CallbackUrl = http://localhost:8080/services/his-core -> .../services/his-core/telerad/auth/token
const (
	partnerAuthTokenPath     = "/telerad/auth/token"
	partnerReadingResultPath = "/telerad/reading-order/actions/result"
)

// tokenExpirySafetyMargin trừ hao vào hạn token đề phòng clock giữa các hệ thống lệch nhau.
const tokenExpirySafetyMargin = 5 * time.Minute

var partnerHttpClient = &http.Client{Timeout: 30 * time.Second}

// mapTokenCache cache token theo từng đối tác (mỗi telerad_partner 1 credential riêng).
var mapTokenCache = make(map[uuid.UUID]partnerToken)
var tokenCacheMutex sync.RWMutex

// ========================= helper dùng chung =========================

func writeLog(httpStatus int, duration time.Duration, url string, requestBody any, responseBody []byte) {
	reqBody := ""
	resBody := ""

	if requestBody != nil {
		if b, err := json.Marshal(requestBody); err == nil {
			reqBody = string(b)
		}
	}
	if len(responseBody) > 0 {
		resBody = string(responseBody)
	}

	logger.Infof("partner callback result: status: %d | latency: %v | url: %s | requestBody: %s | responseBody: %s",
		httpStatus, duration, url, reqBody, resBody)
}

// getToken trả access token còn hạn của đối tác: cache-hit thì dùng lại; chưa có /
// hết hạn thì login his-core (CallbackUrl + PartnerUsername/PartnerPassword), cache lại.
func getToken(partner *entities.TeleradPartnerEntity) (string, error) {
	// fast path: chỉ cần read lock
	tokenCacheMutex.RLock()
	cached, ok := mapTokenCache[partner.Uuid]
	tokenCacheMutex.RUnlock()
	if ok && time.Now().Before(cached.ExpiresAt) {
		return cached.AccessToken, nil
	}

	tokenCacheMutex.Lock()
	defer tokenCacheMutex.Unlock()

	// double-check sau khi giành write lock
	if cached, ok := mapTokenCache[partner.Uuid]; ok && time.Now().Before(cached.ExpiresAt) {
		return cached.AccessToken, nil
	}

	if partner.CallbackUrl == nil || strings.TrimSpace(*partner.CallbackUrl) == "" ||
		partner.PartnerUsername == nil || partner.PartnerPassword == nil {
		return "", errors.New("đối tác chưa đủ cấu hình callback (url/username/password)")
	}

	baseUrl := strings.TrimRight(*partner.CallbackUrl, "/")
	url := baseUrl + partnerAuthTokenPath

	requestDataJson, err := json.Marshal(partnerRequests.LoginRequest{Username: *partner.PartnerUsername, Password: *partner.PartnerPassword})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestDataJson))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	startCallingTime := time.Now()
	response, err := partnerHttpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	// không log request body của login để tránh lộ password
	writeLog(response.StatusCode, time.Since(startCallingTime), url, nil, data)

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("đăng nhập đối tác lỗi (HTTP %d): %s", response.StatusCode, string(data))
	}

	var loginResponse partnerResponses.LoginResponse
	if err := json.Unmarshal(data, &loginResponse); err != nil {
		return "", fmt.Errorf("không đọc được token đối tác: %s", err.Error())
	}
	if loginResponse.AccessToken == "" {
		return "", errors.New("đối tác không trả về token")
	}

	mapTokenCache[partner.Uuid] = partnerToken{
		AccessToken: loginResponse.AccessToken,
		ExpiresAt:   time.Now().Add(time.Duration(loginResponse.ExpiresIn)*time.Second - tokenExpirySafetyMargin),
	}
	return loginResponse.AccessToken, nil
}

// ========================= api đối tác =========================

// SendReadingResultToPartner đẩy kết quả 1 ca đã DUYỆT về đối tác (his-core) qua
// callback. Build payload từ entity ca đọc + tên bác sĩ duyệt; token được getToken
// tự lo (login khi cache hết hạn).
func SendReadingResultToPartner(
	ctx context.Context,
	partner *entities.TeleradPartnerEntity,
	readingOrder *entities.TeleradReadingOrderEntity,
	approvedByName string,
) error {
	if partner.CallbackUrl == nil || strings.TrimSpace(*partner.CallbackUrl) == "" {
		return errors.New("đối tác chưa cấu hình callback URL")
	}
	url := strings.TrimRight(*partner.CallbackUrl, "/") + partnerReadingResultPath

	resultedAt := time.Now().Format(time.RFC3339)
	if readingOrder.ApprovedAt != nil {
		resultedAt = readingOrder.ApprovedAt.Format(time.RFC3339)
	}

	payload := partnerRequests.SendReadingResultRequest{
		TeleradReadingOrderUuid: readingOrder.Uuid.String(),
		OrderId:                 readingOrder.OrderId,
		OrderCode:               readingOrder.OrderCode,
		OrderItemId:             readingOrder.OrderItemId,
		OrderItemCode:           readingOrder.OrderItemCode,
		StudyInstanceUid:        readingOrder.StudyInstanceUid,
		Status:                  readingOrder.Status,
		ResultInHtml:            readingOrder.ResultInHtml,
		ResultInText:            readingOrder.ResultInText,
		ResultedAt:              resultedAt,
		ApprovedByName:          approvedByName,
	}

	token, err := getToken(partner)
	if err != nil {
		return err
	}

	requestDataJson, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(requestDataJson))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	startCallingTime := time.Now()
	response, err := partnerHttpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	writeLog(response.StatusCode, time.Since(startCallingTime), url, payload, data)

	var baseResponse partnerResponses.BaseResponse
	_ = json.Unmarshal(data, &baseResponse)

	if response.StatusCode != http.StatusOK || baseResponse.Code != http.StatusOK {
		msg := baseResponse.Message
		if msg == "" {
			msg = string(data)
		}
		return fmt.Errorf("HTTP %d: %s", response.StatusCode, msg)
	}
	return nil
}
