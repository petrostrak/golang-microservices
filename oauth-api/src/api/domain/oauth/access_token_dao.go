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
