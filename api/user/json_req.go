package user

type registerReq struct {
	Username        string `json:"username" binding:"required" extensions:"x-order=1"`
	Email           string `json:"email" binding:"required" extensions:"x-order=2"`
	Phone           string `json:"phone" binding:"required" extensions:"x-order=3"`
	Password        string `json:"password" binding:"required" extensions:"x-order=4"`
	ConfirmPassword string `json:"confirmPassword" binding:"required" extensions:"x-order=5"`
}

type loginReq struct {
	Username string `json:"username" binding:"required" extensions:"x-order=1"`
	Password string `json:"password" binding:"required" extensions:"x-order=4"`
}
