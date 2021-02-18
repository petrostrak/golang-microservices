package services

import (
	"golang-microservices/oauth-api/src/api/domain/oauth"
	"golang-microservices/src/api/utils/errors"
	"time"
)

type oauthService struct{}

type oauthServiceInterface interface {
	CreateAccessToken(req oauth.AccessTokenRequest) (*oauth.AccessToken, errors.ApiError)
	GetAccessToken(accessToken string) (*oauth.AccessToken, errors.ApiError)
}

var (
	// OauthService is defined as an interface
	OauthService oauthServiceInterface
)

func init() {
	// dependency injection
	// instantiating the service as a pointer to
	// oauthService struct
	OauthService = &oauthService{}
}

// CreateAccessToken creates a token everytime the user is asking for
func (s *oauthService) CreateAccessToken(req oauth.AccessTokenRequest) (*oauth.AccessToken, errors.ApiError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	user, err := oauth.GetUserByUsernamAndPassWord(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	token := oauth.AccessToken{
		UserID:  user.ID,
		Expires: time.Now().UTC().Add(24 * time.Hour).Unix(),
	}

	if err := token.Save(); err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *oauthService) GetAccessToken(accessToken string) (*oauth.AccessToken, errors.ApiError) {
	token, err := oauth.GetAccessTokenByToken(accessToken)
	if err != nil {
		return nil, err
	}

	return token, err
}
