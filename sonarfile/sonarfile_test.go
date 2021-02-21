package sonarfile_test

import (
	"testing"

	"github.com/kgoins/sonar-client/sonarfile"
	"github.com/stretchr/testify/require"
)

func TestBuildSonarFile(t *testing.T) {
	rq := require.New(t)

	fileName := "2021-01-06-1609894956-http_get_9200.csv.gz"
	validSonarFile := sonarfile.SonarFile{
		Date:        "2021-01-06",
		Epoch:       1609894956,
		ServiceName: "http_get",
		Port:        9200,
		Ext:         "csv.gz",
	}

	sonarFile, err := sonarfile.BuildSonarFile(fileName)

	rq.NoError(err)
	rq.Equal(sonarFile, validSonarFile)
}
