package pyro

import (
	"io"
	"os"
	"strings"
)

func parseRoute(r io.Reader) Route {
	// things
	return Route{
		Url:    "url",
		Status: "status",
	}
}

func ParseSuite(r io.Reader) Suite {
	s := Suite{}
}
