package main

import (
	"net/http"
	"os"
	"time"

	"github.com/thesephist/pyro/pkg/pyro"
)

func main() {
	c := pyro.Client{
		Http: http.Client{
			Timeout: 5 * time.Second,
		},
	}

	suite := pyro.NewSuite(
		"https://google.com",
		"https://hackclub.com",
		"https://hackclub.com/nonexistent.html",
		"http://broken-website-tester-324.co.kr",
	)
	suite.AddRoute(pyro.Route{
		Url:    "https://yahoo.com/some_nonexistent_page_243.html",
		Status: 404,
	})

	if !c.RunSuite(suite) {
		os.Exit(1)
	}
}
