package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Hello(c echo.Context) error {
	r := c.Request()

	return c.String(http.StatusOK, fmt.Sprintf("Hello, World!: %s", r.URL.Path))
}

func Health(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
