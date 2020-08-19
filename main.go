package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	p := os.Getenv("PORT")
	if p == "" { // set a default for the port if none is provided
		p = "1323"
	}
	h := fmt.Sprintf(":%s", p)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(h))
}
