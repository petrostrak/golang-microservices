package oauth

import (
	"fmt"
	"golang-microservices/src/api/utils/errors"
)

var (
	tokens = make(map[string]*AccessToken, 0)
)

func (at *AccessToken) Save() errors.ApiError {
	at.AccessToken = fmt.Sprintf("%d", at.UserID)
	tokens[at.AccessToken] = at
	return nil
}

func GetAccessTokenByToken(AccessToken string) (*AccessToken, errors.ApiError) {
	token := tokens[AccessToken]
	if token == nil || token.IsExpired() {
		return nil, errors.NewNotFoundError("no access token found")
	}
	return token, nil
}
