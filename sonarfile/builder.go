package sonarfile

import (
	"errors"
	"strconv"
	"strings"
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

// SplitServiceName will extract the base service
// and port from a Sonar service name
func SplitServiceName(fullServiceName string) (string, int, error) {
	serviceArr := strings.Split(fullServiceName, "_")

	serviceName := strings.Join(serviceArr[:len(serviceArr)-1], "_")
	port, err := strconv.Atoi(serviceArr[len(serviceArr)-1])

	return serviceName, port, err
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

	serviceName, port, err := SplitServiceName(baseWithoutExt)
	if err != nil {
		return SonarFile{}, err
	}

	return SonarFile{
		Date:        date,
		Epoch:       epoch,
		ServiceName: serviceName,
		Port:        port,
		Ext:         ext,
	}, nil
}
