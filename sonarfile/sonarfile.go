package sonarfile

import (
	"fmt"
	"net/url"
	"path"
	"strconv"

	"github.com/kgoins/sonar-client/sonarservice"
)

// SonarFile is the file naming construct used in Sonar data
type SonarFile struct {
	Date  string
	Epoch int64
	Ext   string
	sonarservice.SonarService
}

func (f SonarFile) GetBaseName() string {
	return f.SonarService.ToString()
}

func (f SonarFile) GetFullFilename() string {
	epochStr := strconv.FormatInt(f.Epoch, 10)
	baseServiceName := f.GetBaseName()

	return fmt.Sprintf(
		"%s-%s-%s.%s",
		f.Date,
		epochStr,
		baseServiceName,
		f.Ext,
	)
}

// GetFileURL constructs the full URL path to a given sonar data file
func (f SonarFile) GetFileURL(baseURL string, studyID string) string {
	targetFilename := f.GetFullFilename()

	targetURL, _ := url.Parse(baseURL)
	targetURL.Path = path.Join(targetURL.Path, studyID, targetFilename)

	return targetURL.String()
}
