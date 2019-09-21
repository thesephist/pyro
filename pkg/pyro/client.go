package pyro

import (
	"fmt"
	"net/http"
	goUrl "net/url"

	"github.com/logrusorgru/aurora"
)

type Client struct {
	Http http.Client
}

func (c Client) Log(msg string) {
	fmt.Printf("%s\n", msg)
}

func (c Client) Ok(url string) bool {
	return c.Check(url, 200)
}

func (c Client) Check(url string, status int) bool {
	return c.Status(url) == status
}

func (c Client) Status(url string) int {
	resp, err := c.Http.Get(url)
	if err != nil {
		if urlErr, isUrlErr := err.(*goUrl.Error); isUrlErr {
			if urlErr.Timeout() {
				c.Log(aurora.Sprintf("%s\t%s",
					aurora.Red("timeout"),
					url,
				))
			} else {
				c.Log(aurora.Sprintf("%s\t%s\n  - %s",
					aurora.BgRed("url error"),
					url,
					urlErr.Error(),
				))
			}
		}
		return -1
	}

	resp.Body.Close()
	return resp.StatusCode
}

func (c Client) RunSuite(s Suite) bool {
	failed := 0
	for _, r := range s.Routes {
		statusCode := c.Status(r.Url)
		if statusCode == r.Status {
			c.Log(aurora.Sprintf("%s\t%s",
				aurora.Green(fmt.Sprintf("pass [%d]", r.Status)),
				r.Url,
			))
		} else {
			if statusCode != -1 {
				c.Log(aurora.Sprintf("%s\t%s",
					aurora.Red(fmt.Sprintf("error [%d]", r.Status)),
					r.Url,
				))
			}

			failed++
		}
	}

	if failed == 0 {
		fmt.Printf("\n%s\n", aurora.Green("All routes passed!"))
		return true
	} else {
		fmt.Printf("\n%s\n", aurora.Sprintf(aurora.Red("%d routes failed."), failed))
		return false
	}
}
