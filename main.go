package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// Audit is a simple record that something happened
type Audit struct {
	Utc  string `json:"utc" xml:"utc"`
	Path string `json:"path" xml:"path"`
}

func main() {
	e := echo.New()
	e.GET("/one", func(c echo.Context) error {
		a := &Audit{
			Utc:  time.Time.Format(time.Now(), "2006-01-02 15:04:05.000"),
			Path: "/one",
		}
		return c.JSON(http.StatusOK, a)
	})
	e.GET("/two", func(c echo.Context) error {
		a := &Audit{
			Utc:  time.Time.Format(time.Now(), "2006-01-02 15:04:05.000"),
			Path: "/two",
		}
		return c.JSON(http.StatusOK, a)
	})
	e.Logger.Fatal(e.Start(":8181"))
}
