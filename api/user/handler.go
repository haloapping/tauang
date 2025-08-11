package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService userService
}

func NewUserHandler(s userService) *userHandler {
	return &userHandler{
		userService: s,
	}
}

// Register User godoc
//
//	@Summary		Register user
//	@Description	Register user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		registerReq	true	"Register User Request"
//	@Success		200		{object}	registerResp
//	@Router			/users/register [post]
func (h userHandler) register(c echo.Context) error {
	var reqBody registerReq
	err := c.Bind(&reqBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	errValidation := registerValidation(reqBody)
	if len(errValidation) > 0 {
		return c.JSON(http.StatusBadRequest, errValidation)
	}

	ur, err := h.userService.register(c, reqBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	resp := registerResp{
		Message: "user is registered",
		Data: user{
			Username:  ur.Username,
			Email:     ur.Email,
			Phone:     ur.Phone,
			CreatedAt: ur.CreatedAt,
			UpdatedAt: ur.UpdatedAt,
		},
	}

	return c.JSON(http.StatusCreated, resp)
}

// Login User godoc
//
//	@Summary		Login user
//	@Description	Login user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		loginReq	true	"Login User Request"
//	@Success		200		{object}	loginResp
//	@Router			/users/login [post]
func (h userHandler) login(c echo.Context) error {
	var reqBody loginReq
	err := c.Bind(&reqBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	errValidation := loginValidation(reqBody)
	if len(errValidation) > 0 {
		return c.JSON(http.StatusBadRequest, errValidation)
	}

	token, err := h.userService.login(c, reqBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	resp := loginResp{
		Token: token,
	}

	return c.JSON(http.StatusOK, resp)
}
