package model

import "time"

type (
	CustomerLoginResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiredAt    string `json:"expired_at"`
	}
	CustomersDetail struct {
		Id        int       `json:"id"`
		Email     string    `json:"email"`
		Phone     string    `json:"phone"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	}
)
