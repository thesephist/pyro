package pyro

import (
	"fmt"
	"net/http"
	goUrl "net/url"
)

type Client struct {
	Name string
	Http http.Client
}

func (c Client) Log(msg string) {
	fmt.Printf("%s: %s\n", c.Name, msg)
}

func (c Client) Ok(url string) bool {
	return c.Check(url, 200)
}

func (c Client) Check(url string, status int) bool {
	resp, err := c.Http.Get(url)
	if err != nil {
		if urlErr, isUrlErr := err.(*goUrl.Error); isUrlErr {
			if urlErr.Timeout() {
				c.Log(fmt.Sprintf("timeout\t%s", url))
			} else {
				c.Log(fmt.Sprintf("urlerr\t%s\n  - %s", url, urlErr.Error()))
			}
		}
		return false
	}

	resp.Body.Close()
	return resp.StatusCode == status
}

func (c Client) RunSuite(s Suite) bool {
	allPassed := true
	for _, r := range s.Routes {
		if c.Check(r.Url, r.Status) {
			c.Log(fmt.Sprintf("ok\t%s", r.Url))
		} else {
			allPassed = false
		}
	}

	return allPassed
}
