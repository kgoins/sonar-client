package cmd

import (
	"fmt"
	"log"

	"github.com/kgoins/sonar-client/pageclient"
	"github.com/spf13/cobra"
)

var studiesCmd = &cobra.Command{
	Use:   "studies",
	Short: "Lists the current studies available from Rapid7's OpenData platform",
	Run: func(cmd *cobra.Command, args []string) {
		studies, err := pageclient.ListStudies()
		if err != nil {
			log.Fatalln(err.Error())
		}

		for _, s := range studies {
			fmt.Println(s)
		}
	},
}

func init() {
	rootCmd.AddCommand(studiesCmd)
}
