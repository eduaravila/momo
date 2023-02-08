package types

import "time"

type (
	TokenBodyRequest struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		Code         string `json:"code"`
		GrantType    string `json:"grant_type"`
		RedirectURI  string `json:"redirect_uri"`
	}

	TokenResponse struct {
		AccessToken  string   `json:"access_token"`
		RefreshToken string   `json:"refresh_token"`
		ExpiresIn    int      `json:"expires_in"`
		TokenType    string   `json:"token_type"`
		Scope        []string `json:"scope"`
	}
)

type UserinfoRespose struct {
	Aud              string    `json:"aud"`
	Exp              int64     `json:"exp"`
	Iat              int64     `json:"iat"`
	Iss              string    `json:"iss"`
	Sub              string    `json:"sub"`
	Email            string    `json:"email"`
	EmailVerified    bool      `json:"email_verified"`
	Picture          string    `json:"picture"`
	PreferedUsername string    `json:"preferred_username"`
	UpdatedAt        time.Time `json:"updated_at"`
}
