package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Verbose bool
var Json bool

var rootCmd = &cobra.Command{
	Use:   "asite",
	Short: "Site analysis CLI",
	Long:  `Site analysis CLI using Google's PageSpeed API.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error ocurred!")
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Outputting relevant information")
}
