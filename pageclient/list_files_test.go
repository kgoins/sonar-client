package pageclient_test

import (
	"testing"

	"github.com/kgoins/sonar-client/pageclient"
	"github.com/stretchr/testify/require"
)

func TestListFiles(t *testing.T) {
	rq := require.New(t)

	fileNames, err := pageclient.ListFiles("sonar.http")
	rq.NoError(err)
	rq.NotEmpty(fileNames)
}
