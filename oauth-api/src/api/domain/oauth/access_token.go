package oauth

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	Expires     int64  `json:"expires"`
}
