package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/haloapping/tauang/api/user"
	"github.com/haloapping/tauang/api/wallet"
	"github.com/haloapping/tauang/db"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//	@title			tauang API
//	@version		1.0
//	@description	tauang API.

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@securityDefinitions.apikey	JWT
//	@in							header
//	@name						Authorization

// @host		localhost:3000
// @BasePath	/
func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
	pool := db.NewConnection(connString)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	userRepo := user.NewUserRepository(pool)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)
	user.Router(e.Group("/users"), userHandler)

	walletRepo := wallet.NewWalletRepository(pool)
	walletService := wallet.NewWalletService(walletRepo)
	walletHandler := wallet.NewWalletHandler(walletService)
	wallet.Router(e.Group("/wallets"), walletHandler)

	e.GET("/", func(c echo.Context) error {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "tauang API",
			},
			DarkMode:   true,
			Theme:      scalar.ThemeDeepSpace,
			Layout:     scalar.LayoutModern,
			HideModels: true,
		})

		if err != nil {
			fmt.Printf("%v", err)
			return c.String(http.StatusInternalServerError, "Failed to generate API reference")
		}

		return c.HTML(http.StatusOK, htmlContent)
	})

	fmt.Printf("Starting web server on port :3000")
	e.Logger.Fatal(e.Start(":3000"))
}
