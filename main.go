package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// Audit is a simple record that something happened
type Audit struct {
	UTC           string `json:"utc" xml:"utc"`
	URI           string `json:"path" xml:"path"`
	HeaderBodyMD5 string `json:"header_body_md5" xml:"header_body_md5"`
	HeaderStatus  string `json:"header_status" xml:"header_status"`
}

func main() {
	e := echo.New()
	e.GET("/one/*", serveJSON)
	e.GET("/two/*", serveJSON)
	e.Logger.Fatal(e.Start(":8181"))
}
func serveJSON(c echo.Context) error {
	r := c.Request()
	b := r.Header.Get("X-Body-MD5")
	s := r.Header.Get("X-Status")
	u := r.RequestURI
	a := &Audit{
		UTC:           time.Time.Format(time.Now(), "2006-01-02 15:04:05.000"),
		URI:           u,
		HeaderBodyMD5: b,
		HeaderStatus:  s,
	}
	return c.JSON(http.StatusOK, a)
}
