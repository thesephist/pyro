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

func (c Client) Check(url string, status int) bool {
	respStatus := c.Status(url)
	if respStatus == status {
		fmt.Println(aurora.Green(fmt.Sprintf("pass [%d]", respStatus)))
		return true
	} else {
		fmt.Println(aurora.Red(fmt.Sprintf("error [-> %d != %d]", respStatus, status)))
		return false
	}
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
		respStatus := c.Status(r.Url)
		if respStatus == r.Status {
			c.Log(aurora.Sprintf("%s\t%s",
				aurora.Green(fmt.Sprintf("pass [%d]", respStatus)),
				r.Url,
			))
		} else {
			if respStatus != -1 {
				c.Log(aurora.Sprintf("%s\t%s",
					aurora.Red(fmt.Sprintf("error [%d]", respStatus)),
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
