package sonarfile_test

import (
	"testing"

	"github.com/kgoins/sonar-client/sonarfile"
	"github.com/stretchr/testify/require"
)

func TestBuildSonarFile(t *testing.T) {
	rq := require.New(t)
	testCases := make(map[string]sonarfile.SonarFile)

	testCases["2021-01-06-1609894956-http_get_9200.csv.gz"] = sonarfile.SonarFile{
		Date:        "2021-01-06",
		Epoch:       1609894956,
		ServiceName: "http_get",
		Port:        9200,
		Ext:         "csv.gz",
	}
	testCases["2021-01-04-1609768117-gtp-c_2123.csv.gz"] = sonarfile.SonarFile{
		Date:        "2021-01-04",
		Epoch:       1609768117,
		ServiceName: "gtp-c",
		Port:        2123,
		Ext:         "csv.gz",
	}
	testCases["2021-01-04-1609759821-udp_chargen_19.csv.gz"] = sonarfile.SonarFile{
		Date:        "2021-01-04",
		Epoch:       1609759821,
		ServiceName: "chargen",
		Port:        19,
		Ext:         "csv.gz",
	}

	for fileName, expectedSonarFile := range testCases {
		sonarFile, err := sonarfile.BuildSonarFile(fileName)
		rq.NoError(err)
		rq.Equal(expectedSonarFile, sonarFile)
	}
}
