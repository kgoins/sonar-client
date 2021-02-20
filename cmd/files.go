package cmd

import (
	"fmt"
	"log"

	"github.com/kgoins/sonar-client/pageclient"
	"github.com/spf13/cobra"
)

var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "Lists the current files available for download from a given study",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		studyID := args[0]
		files, err := pageclient.ListFiles(studyID)
		if err != nil {
			log.Fatalln(err.Error())
		}

		for _, f := range files {
			fmt.Println(f)
		}
	},
}

func init() {
	rootCmd.AddCommand(filesCmd)
}
