package main

import (
	"net/http"
	"os"
	"time"

	"github.com/thesephist/pyro/pkg/pyro"
)

func main() {
	c := pyro.Client{
		Name: "Pyro",
		Http: http.Client{
			Timeout: 5 * time.Second,
		},
	}

	suite := pyro.NewSuite(
		"https://google.com",
		"https://hackclub.com",
		"http://broken-website-tester-324.co.kr",
	)

	if !c.RunSuite(suite) {
		os.Exit(1)
	}
}
