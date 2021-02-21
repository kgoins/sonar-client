package sonarfile

import (
	"errors"
	"strconv"
	"strings"

	"github.com/kgoins/sonar-client/sonarservice"
)

// extractExtension splits a filename into the name itself
// and the extension
func extractExtension(filename string) (string, string) {
	nameSplit := strings.SplitN(filename, ".", 2)
	if len(nameSplit) == 0 {
		return "", ""
	}

	if len(nameSplit) < 2 {
		return nameSplit[0], ""
	}

	return nameSplit[0], nameSplit[1]
}

// BuildSonarFile constructs a SonarFile from a string
// Example: 2021-01-06-1609894956-http_get_9200.csv.gz
func BuildSonarFile(filename string) (SonarFile, error) {
	fileArray := strings.Split(filename, "-")
	if len(fileArray) != 5 {
		return SonarFile{}, errors.New("Invalid format")
	}

	date := strings.Join(fileArray[:3], "-")
	epoch, err := strconv.ParseInt(fileArray[3], 10, 64)
	if err != nil {
		return SonarFile{}, err
	}

	baseName := fileArray[4]
	baseWithoutExt, ext := extractExtension(baseName)

	service, err := sonarservice.NewSonarService(baseWithoutExt)
	if err != nil {
		return SonarFile{}, err
	}

	return SonarFile{
		Date:        date,
		Epoch:       epoch,
		ServiceName: service.Name,
		Port:        service.Port,
		Ext:         ext,
	}, nil
}
