package pyro

import (
	"io"
	_ "os"
	_ "strings"
)

func parseRoute(r io.Reader) Route {
	// things
	return Route{
		Url:    "url",
		Status: 200,
	}
}

func ParseSuite(r io.Reader) Suite {
	s := Suite{}
	return s
}
