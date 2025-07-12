package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var singleCmd = &cobra.Command{
	Use:   "url [arg]",
	Short: "Single URL to analyse",
	Long:  `Single url to be passed for analysis with Google's PageSpeed API`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("more than one argument provided")
		} else if len(args) < 1 {
			return errors.New("no argument is provided")
		}

		url := args[0]
		if !isValidUrl(url) {
			return fmt.Errorf("%s is not a valid URL format", url)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(singleCmd)
}
