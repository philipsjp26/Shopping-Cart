package entity

import "time"

type AccessToken struct {
	Id           uint       `json:"id"`
	CustomerID   uint       `json:"customer_id"`
	AccessToken  string     `json:"access_token"`
	RefreshToken string     `json:"refresh_token"`
	RevokedAt    *time.Time `json:"revoked_at"`
	CreatedAt    *time.Time `json:"created_at"`
}
