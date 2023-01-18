package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello() string {
	return "Hello, world."
}

func main() {
	fmt.Print(ViperEnvVariable("HELLO_WORLD"))
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, Hello())
	})
	e.Logger.Fatal(e.Start(":8080"))
}
