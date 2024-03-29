package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/thesephist/pyro/pkg/pyro"
)

var checkCmd = &cobra.Command{
	Use:   "check <url> <status>",
	Short: "Check a single route against an expected status code",
	Long: `pyro check confirms that a given URL returns a valid HTTP
response, and that the response status code matches the
expected status code. If one is not provided, a status code
of OK 200 is assumed.`,
	Run: func(cmd *cobra.Command, args []string) {
		url := ""
		status := 200

		if len(args) == 1 {
			url = args[0]
		} else if len(args) == 2 {
			url = args[0]

			statusNum, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Printf("error: invalid status code %s\n", args[1])
				return
			} else {
				status = statusNum
			}
		} else {
			fmt.Println("error: invalid syntax for pyro check <url> <status>")
			return
		}

		c := pyro.Client{}
		c.Check(url, status)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
