package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var multipleCmd = &cobra.Command{
	Use:   "file [arg]",
	Short: "Pass text file for list of URLs",
	Long:  `Passing a text file to run Google PageSpeed API on list of URLs`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("more than one argument provided")
		} else if len(args) < 1 {
			return errors.New("no argument is provided")
		}

		fp := args[0]
		info, err := os.Stat(fp)
		if err != nil {
			if os.IsNotExist(err) {
				return errors.New("file/directory does not exist")
			}
		}
		if info.IsDir() {
			return fmt.Errorf("%s is a directory not a file", fp)
		}
		fExtension := strings.Split(info.Name(), ".")

		if len(fExtension) != 2 {
			if len(fExtension) > 2 {
				return fmt.Errorf("%s is a file with a compound extension and not .txt", fp)
			} else {
				return fmt.Errorf("%s is a file with no extension", fp)
			}
		}
		if fExtension[len(fExtension)-1] != "txt" {
			return fmt.Errorf("%s is a file with a different extension than .txt", fp)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("file called")
	},
}

func init() {
	rootCmd.AddCommand(multipleCmd)
}
