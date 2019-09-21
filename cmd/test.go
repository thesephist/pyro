package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/thesephist/pyro/pkg/pyro"
)

type SpecParseErr struct {
	message string
}

func (spe SpecParseErr) Error() string {
	return spe.message
}

func suiteFromSpec(specFile io.Reader) (pyro.Suite, error) {
	suite := pyro.Suite{}

	spec, err := ioutil.ReadAll(specFile)
	if err != nil {
		return suite, SpecParseErr{"could not read spec file"}
	}
	specLines := strings.Split(string(spec), "\n")

	for _, line := range specLines {
		if len(line) == 0 {
			continue
		}

		splitLine := strings.Split(line, " ")
		if len(splitLine) == 1 {
			suite.AddRoute(pyro.Route{
				Url:    splitLine[0],
				Status: 200,
			})
		} else if len(splitLine) >= 2 {
			status, err := strconv.Atoi(splitLine[1])
			if err != nil {
				return suite, SpecParseErr{fmt.Sprintf("could not parse line: '%s'", line)}
			}

			suite.AddRoute(pyro.Route{
				Url:    splitLine[0],
				Status: status,
			})
		}
	}

	return suite, nil
}

func testFromSuite(suite pyro.Suite) {
	c := pyro.Client{
		Http: http.Client{
			Timeout: 5 * time.Second,
		},
	}

	if !c.RunSuite(suite) {
		os.Exit(1)
	}
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Pyro test",
	Long:  "Pyro long test suite",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("error: invalid syntax for pyro test <file>")
			return
		}

		specFile, err := os.Open(args[0])
		defer specFile.Close()
		if err != nil {
			fmt.Printf("error: could not read test spec file %s\n", args[0])
			return
		}

		suite, err := suiteFromSpec(specFile)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		testFromSuite(suite)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
