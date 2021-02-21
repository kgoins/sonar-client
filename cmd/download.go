package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/kgoins/sonar-client/pageclient"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download studyID service",
	Short: "Downloads the specified sonar dataset file",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		studyID := args[0]
		service := args[1]

		cwd, _ := os.Getwd()
		outFilePath := filepath.Join(
			cwd,
			service,
		)

		outfile, _ := cmd.Flags().GetString("outfile")
		if outfile != "" {
			outFilePath = outfile
		}

		err := pageclient.DownloadServiceData(studyID, service, outFilePath)
		if err != nil {
			log.Fatalln(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().StringP(
		"outfile",
		"o",
		"",
		"Path and name of the file that data should be downloaded to. Defaults to the service name downloaded to the current directory.",
	)
}
