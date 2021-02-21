package sonarfile_test

import (
	"testing"

	"github.com/kgoins/sonar-client/sonarfile"
	"github.com/kgoins/sonar-client/sonarservice"
	"github.com/stretchr/testify/require"
)

func TestBuildSonarFile(t *testing.T) {
	rq := require.New(t)
	testCases := make(map[string]sonarfile.SonarFile)

	testCases["2021-01-06-1609894956-http_get_9200.csv.gz"] = sonarfile.SonarFile{
		Date:  "2021-01-06",
		Epoch: 1609894956,
		Ext:   "csv.gz",
		SonarService: sonarservice.SonarService{
			Name: "http_get",
			Port: 9200,
		},
	}
	testCases["2021-01-04-1609768117-gtp-c_2123.csv.gz"] = sonarfile.SonarFile{
		Date:  "2021-01-04",
		Epoch: 1609768117,
		Ext:   "csv.gz",
		SonarService: sonarservice.SonarService{
			Name: "gtp-c",
			Port: 2123,
		},
	}
	testCases["2021-01-04-1609759821-udp_chargen_19.csv.gz"] = sonarfile.SonarFile{
		Date:  "2021-01-04",
		Epoch: 1609759821,
		Ext:   "csv.gz",
		SonarService: sonarservice.SonarService{
			Name:      "chargen",
			Port:      19,
			Transport: sonarservice.UDP,
		},
	}

	for fileName, expectedSonarFile := range testCases {
		sonarFile, err := sonarfile.BuildSonarFile(fileName)
		rq.NoError(err)
		rq.Equal(expectedSonarFile, sonarFile)
	}
}

func TestGetBaseName(t *testing.T) {
	rq := require.New(t)

	testFile := sonarfile.SonarFile{
		Date:  "2021-01-04",
		Epoch: 1609759821,
		Ext:   "csv.gz",
		SonarService: sonarservice.SonarService{
			Name:      "chargen",
			Port:      19,
			Transport: sonarservice.UDP,
		},
	}

	expected := "udp_chargen_19"
	filename := testFile.GetBaseName()

	rq.Equal(expected, filename)
}

func TestGetFullFilename(t *testing.T) {
	rq := require.New(t)

	testFile := sonarfile.SonarFile{
		Date:  "2021-01-04",
		Epoch: 1609759821,
		Ext:   "csv.gz",
		SonarService: sonarservice.SonarService{
			Name:      "chargen",
			Port:      19,
			Transport: sonarservice.UDP,
		},
	}

	expected := "2021-01-04-1609759821-udp_chargen_19.csv.gz"
	filename := testFile.GetFullFilename()

	rq.Equal(expected, filename)
}
