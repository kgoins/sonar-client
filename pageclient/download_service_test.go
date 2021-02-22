package pageclient_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/kgoins/sonar-client/pageclient"
	"github.com/kgoins/sonar-client/sonarservice"
	"github.com/stretchr/testify/require"
)

func getTestDataDir() string {
	_, mypath, _, _ := runtime.Caller(0)
	myparent := filepath.Dir(mypath)
	return filepath.Join(myparent, "testdata")
}

func TestDownloadService(t *testing.T) {
	rq := require.New(t)

	studyID := "sonar.udp"
	service := sonarservice.SonarService{
		Name:      "chargen",
		Port:      19,
		Transport: sonarservice.UDP,
	}

	testOut := filepath.Join(getTestDataDir(), "out.csv.gz")
	os.Remove(testOut)

	err := pageclient.DownloadServiceData(studyID, service, testOut)
	rq.NoError(err)
	rq.FileExists(testOut)

	os.Remove(testOut)
}
