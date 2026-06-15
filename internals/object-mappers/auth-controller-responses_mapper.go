package objectMappers

import (
	"telerad-core-module/internals/constants"
	authControllerResponses "telerad-core-module/internals/responses/auth-controller_responses"
)

func ToGetTokenResponse(accessToken string, expiresIn int64) authControllerResponses.GetTokenResponse {
	return authControllerResponses.GetTokenResponse{
		AccessToken: accessToken,
		TokenType:   constants.JWT_TOKEN_TYPE,
		ExpiresIn:   expiresIn,
	}
}
