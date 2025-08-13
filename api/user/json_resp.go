package user

import "github.com/guregu/null/v6"

type user struct {
	Username  string    `json:"username" extensions:"x-order=1"`
	Email     string    `json:"email" extensions:"x-order=2"`
	Phone     string    `json:"phone" extensions:"x-order=3"`
	CreatedAt null.Time `json:"createdAt" extensions:"x-order=5"`
	UpdatedAt null.Time `json:"updatedAt" extensions:"x-order=6"`
}

type registerResp struct {
	Message string `json:"message" extensions:"x-order=1"`
	Data    user   `json:"data" extensions:"x-order=2"`
}

type loginResp struct {
	Token string `json:"token" extensions:"x-order=1"`
}
