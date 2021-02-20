package pageclient

import "github.com/kgoins/sonar-client/sonarfile"

// ListServices returns the base service name of all
// service / port combos scanned by the study.
// This list of names can be used to request the latest
// version of a study's data without knowing the date
// scanned.
func ListServices(studyID string) ([]string, error) {
	files, err := ListFiles(studyID)
	if err != nil {
		return nil, err
	}

	services := make([]string, 0, len(files))
	for _, fileName := range files {
		sonarFile, err := sonarfile.BuildSonarFile(fileName)
		if err != nil {
			continue
		}

		services = append(services, sonarFile.BuildBaseName())
	}

	return services, nil
}
