package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("Pass the port as the only argument")
	}
	e := echo.New()
	e.GET("/*", serveJSON)
	p := fmt.Sprintf(":%s", args[0])
	e.Logger.Fatal(e.Start(p))
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
	log.Println(u)
	return c.JSON(http.StatusOK, a)
}
