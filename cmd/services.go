package cmd

import (
	"fmt"
	"log"

	"github.com/kgoins/sonar-client/pageclient"
	"github.com/spf13/cobra"
)

var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Lists the current services with scan data available for a given study id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		studyID := args[0]
		services, err := pageclient.ListServices(studyID)
		if err != nil {
			log.Fatalln(err.Error())
		}

		for _, s := range services {
			fmt.Println(s)
		}
	},
}

func init() {
	rootCmd.AddCommand(servicesCmd)
}
