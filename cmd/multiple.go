package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var multipleCmd = &cobra.Command{
	Use:   "file [arg]",
	Short: "Pass text file for list of URLs",
	Long:  `Passing a text file to run Google PageSpeed API on list of URLs`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("file called")
	},
}

func init() {
	rootCmd.AddCommand(multipleCmd)
}
