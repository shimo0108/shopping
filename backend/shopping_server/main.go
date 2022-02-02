package main

import (
	"net/http"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	router := NewRouter()

	router.Start(":9999")
}

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	initRouting(e)

	return e
}

func initRouting(e *echo.Echo) {
	e.GET("/test", testHandler)
}

func testHandler(c echo.Context) error {
	return c.String(http.StatusOK, "health check ok")
}
