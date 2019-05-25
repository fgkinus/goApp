package src

import (
	"net/http"

	"github.com/labstack/echo"
)

// this is the applications default  route
func defaultRoute(c echo.Context) error {
	type message struct {
		Message string `json:"message"`
		Version string `json:"version"`
	}
	m := message{
		Message: "Welcome to the file finder. If you are seeing this, Everything is running topsy turvy",
		Version: "1.0",
	}
	return c.JSON(http.StatusOK, m)
}

func RoutesManager(e *echo.Echo) *echo.Echo {
	// default response to app
	e.GET("/", defaultRoute)

	return e
}
