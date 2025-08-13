package wallet

import "github.com/guregu/null/v6"

type wallet struct {
	ID          string    `json:"ID" binding:"required" extensions:"x-order=1"`
	UserID      string    `json:"userID" binding:"required" extensions:"x-order=2"`
	Name        string    `json:"name" binding:"required" extensions:"x-order=3"`
	Description string    `json:"description" binding:"required" extensions:"x-order=4"`
	Currency    string    `json:"currency" binding:"required" extensions:"x-order=5"`
	CreatedAt   null.Time `json:"createdAt" binding:"required" extensions:"x-order=6"`
	UpdatedAt   null.Time `json:"updatedAt" binding:"required" extensions:"x-order=7"`
}

type singleWalletResp struct {
	Message string `json:"message" binding:"required" extensions:"x-order=1"`
	Data    wallet `json:"data" binding:"required" extensions:"x-order=2"`
}

type multipleWalletResp struct {
	Message string   `json:"message" binding:"required" extensions:"x-order=1"`
	Data    []wallet `json:"data" binding:"required" extensions:"x-order=2"`
}
