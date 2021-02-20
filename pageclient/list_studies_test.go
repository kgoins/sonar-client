package pageclient_test

import (
	"testing"

	"github.com/kgoins/sonar-client/pageclient"
	"github.com/stretchr/testify/require"
)

func TestListStudies(t *testing.T) {
	rq := require.New(t)

	studyNames, err := pageclient.ListStudies()
	rq.NoError(err)
	rq.NotEmpty(studyNames)
}
