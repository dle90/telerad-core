package services

import (
	"context"

	"telerad-core-module/configs"
	baseServices "telerad-core-module/internals/base-services"
	objectMappers "telerad-core-module/internals/object-mappers"
	authControllerRequests "telerad-core-module/internals/requests/auth-controller_requests"
	authControllerResponses "telerad-core-module/internals/responses/auth-controller_responses"
	"telerad-core-module/internals/secure"
	"telerad-core-module/jwtchecker"

	_errorMessages "telerad-core-module/error"

	_error "telerad-core-module/error"

	"github.com/golang-jwt/jwt/v4"
)

var tokenExpireAfter = configs.GetJwtExpiryTime()

// StaffLogin xác thực nhân viên bằng username/password (telerad). Nhân viên phải
// đã có tài khoản (username != nil) và đang active.
func StaffLogin(ctx context.Context, request authControllerRequests.StaffLoginRequest) (*authControllerResponses.GetTokenResponse, *_error.SystemError) {
	account, err := baseServices.GetStaffAccountByUsername(ctx, bunNoTransaction, request.Username)
	if err != nil {
		return nil, _error.New(err)
	} else if account == nil || account.Username == nil || account.PasswordHash == nil {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E001_003)
	} else if !secure.VerifyBcryptPassword(*account.PasswordHash, request.Password) {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E001_003)
	} else if !account.IsActive {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E001_002)
	}

	claims := &jwt.MapClaims{
		secure.JWT_KEY_AUD:       secure.JWT_AUD_TELERAD,
		secure.JWT_KEY_ISS:       secure.JWT_ISS_TELERAD,
		secure.JWT_KEY_UUID:      account.Uuid.String(),
		secure.JWT_KEY_USER_NAME: *account.Username,
		secure.JWT_KEY_TYPE:      secure.JWT_TYPE_STAFF,
		secure.JWT_KEY_ROLES:     account.Roles,
	}

	jwtToken, err := generateToken(claims, tokenExpireAfter)
	if err != nil {
		return nil, _error.New(err)
	}

	result := objectMappers.ToGetTokenResponse(jwtToken, tokenExpireAfter)
	return &result, nil
}

// TeleradPartnerLogin xác thực đối tác telerad bằng credential phía telerad
// (username/password_hash). Đối tác phải đang ở trạng thái ACTIVE.
func TeleradPartnerLogin(ctx context.Context, request authControllerRequests.TeleradPartnerLoginRequest) (*authControllerResponses.GetTokenResponse, *_error.SystemError) {
	partner, err := baseServices.GetTeleradPartnerByUsername(ctx, bunNoTransaction, request.Username)
	if err != nil {
		return nil, _error.New(err)
	} else if partner == nil {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E001_003)
	} else if !secure.VerifyBcryptPassword(partner.PasswordHash, request.Password) {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E001_003)
	} else if !partner.IsActive {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E001_002)
	}

	claims := &jwt.MapClaims{
		secure.JWT_KEY_ISS:       secure.JWT_ISS_TELERAD,
		secure.JWT_KEY_UUID:      partner.Uuid.String(),
		secure.JWT_KEY_USER_NAME: partner.Username,
		secure.JWT_KEY_TYPE:      secure.JWT_TYPE_TELERAD_PARTNER,
	}

	jwtToken, err := generateToken(claims, tokenExpireAfter)
	if err != nil {
		return nil, _error.New(err)
	}

	result := objectMappers.ToGetTokenResponse(jwtToken, tokenExpireAfter)
	return &result, nil
}

func generateToken(claims *jwt.MapClaims, expireAfter int64) (string, error) {
	return jwtchecker.Encode(claims, expireAfter)
}
