package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pyro",
	Short: "Pyro short description",
	Long:  "Pyro long description",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run pyro help to see usage.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
