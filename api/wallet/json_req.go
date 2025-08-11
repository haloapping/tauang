package wallet

type createWalletReq struct {
	UserID      string `json:"userID" binding:"required" extensions:"x-order=1"`
	Name        string `json:"name" binding:"required" extensions:"x-order=2"`
	Description string `json:"description" binding:"required" extensions:"x-order=3"`
	Currency    string `json:"currency" binding:"required" extensions:"x-order=4"`
}

type updateWalletReq struct {
	UserID      string `json:"userID" extensions:"x-order=1"`
	Name        string `json:"name" extensions:"x-order=2"`
	Description string `json:"description" extensions:"x-order=3"`
	Currency    string `json:"currency" extensions:"x-order=4"`
}
