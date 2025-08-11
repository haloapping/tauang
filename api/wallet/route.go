package wallet

import "github.com/labstack/echo/v4"

func Router(g *echo.Group, h *walletHandler) {
	g.POST("", h.create)
}
