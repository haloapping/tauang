package user

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/guregu/null/v6"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repository userRepository
}

func NewUserService(r userRepository) userService {
	return userService{
		repository: r,
	}
}

type userRegister struct {
	ID        string
	Username  string
	Email     string
	Phone     string
	CreatedAt null.Time
	UpdatedAt null.Time
}

func (s userService) register(c echo.Context, req registerReq) (userRegister, error) {
	u, err := s.repository.register(c, req)
	if err != nil {
		return userRegister{}, err
	}

	ur := userRegister{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Phone:     u.Phone,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}

	return ur, nil
}

func (s userService) login(c echo.Context, req loginReq) (string, error) {
	u, err := s.repository.login(c, req)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"username": u.Username,
		"exp":      time.Now().Add(24 * time.Hour),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
