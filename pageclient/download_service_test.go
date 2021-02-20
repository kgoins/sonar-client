package pageclient_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/kgoins/sonar-client/pageclient"
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
	serviceName := "udp_chargen_19"

	testOut := filepath.Join(getTestDataDir(), "out.csv.gz")
	os.Remove(testOut)

	err := pageclient.DownloadServiceData(studyID, serviceName, testOut)
	rq.NoError(err)
	rq.FileExists(testOut)

	os.Remove(testOut)
}
