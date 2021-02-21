package pageclient

import (
	"errors"

	"github.com/cavaliercoder/grab"
	"github.com/kgoins/sonar-client/sonarfile"
)

func findMatchingSonarFile(serviceName string, files []string) (sonarfile.SonarFile, bool) {
	found := false
	var matchingFile sonarfile.SonarFile

	for _, fileName := range files {
		sonarFile, err := sonarfile.BuildSonarFile(fileName)
		if err != nil {
			continue
		}

		if sonarFile.GetBaseName() == serviceName {
			matchingFile = sonarFile
			found = true
		}
	}

	return matchingFile, found
}

// DownloadServiceData will get the datafile for a given service
// inside of a Sonar study and write it to disk.
func DownloadServiceData(studyID, serviceName, outputLocation string) error {
	files, err := ListFiles(studyID)
	if err != nil {
		return err
	}

	matchingFile, found := findMatchingSonarFile(serviceName, files)
	if !found {
		return errors.New("Unable to locate requested service data")
	}

	targetURL := matchingFile.GetFileURL(BaseURL, studyID)

	_, err = grab.Get(outputLocation, targetURL)
	return err
}
