package user

import "github.com/labstack/echo/v4"

func Router(g *echo.Group, h *userHandler) {
	g.POST("/register", h.register)
	g.POST("/login", h.login)
}
