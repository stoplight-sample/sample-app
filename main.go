package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		//AllowHeaders:     []string{echo.HeaderAccessControlAllowHeaders, echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))
}

func initRouting(e *echo.Echo) {
	// 起動確認用API
	e.GET("/", hello)
	e.GET("/bye", bye)
}

// @Tags hello
// @Summary Print Hello World!
// @Description Print Hello World!
// @ID hello
// @Accept  json
// @Produce  json
// @Success 201 {object} string
// @Failure 400 {string} string "errorMessage"
// @Failure 500 {string} string "errorMessage"
// @Router / [get]
func hello(c echo.Context) error {
	log.Println("Hello World!")
	return c.JSON(http.StatusOK, "Hello World!")
}

// @Tags bye
// @Summary Print Bye World!
// @Description Print Bye World!
// @ID bye
// @Accept  json
// @Produce  json
// @Success 201 {object} string
// @Failure 400 {string} string "errorMessage"
// @Failure 500 {string} string "errorMessage"
// @Router /bye [get]
func bye(c echo.Context) error {
	log.Println("Bye World!")
	return c.JSON(http.StatusOK, "Bye World!")
}

func main() {
	e := echo.New()

	initMiddleware(e)
	initRouting(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8055"
		log.Printf("Defaulting to port %s", port)
	}

	if err := e.Start(":" + port); err != nil {
		log.Printf("%v", err)
	}
}
