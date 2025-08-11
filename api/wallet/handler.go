package wallet

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type walletHandler struct {
	walletService walletService
}

func NewWalletHandler(s walletService) *walletHandler {
	return &walletHandler{
		walletService: s,
	}
}

// Create New Wallet godoc
//
//	@Summary		Create new wallet
//	@Description	Create new wallet
//	@Tags			wallets
//	@Accept			json
//	@Produce		json
//	@Param			wallet	body		createWalletReq	true	"Create new wallet request"
//	@Success		200		{object}	singleWalletResp
//	@Router			/wallets [post]
func (h walletHandler) create(c echo.Context) error {
	var reqBody createWalletReq
	err := c.Bind(&reqBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	errValidation := createValidation(reqBody)
	if len(errValidation) > 0 {
		return c.JSON(http.StatusBadRequest, errValidation)
	}

	w, err := h.walletService.create(c, reqBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	resp := singleWalletResp{
		Message: "wallet is created",
		Data:    w,
	}

	return c.JSON(http.StatusCreated, resp)
}
