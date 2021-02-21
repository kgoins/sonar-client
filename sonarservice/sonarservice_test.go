package sonarservice_test

import (
	"testing"

	"github.com/kgoins/sonar-client/sonarservice"
	"github.com/stretchr/testify/require"
)

func TestNewSonarService(t *testing.T) {
	rq := require.New(t)
	testCases := make(map[string]sonarservice.SonarService)

	testCases["tcp_ldap_389"] = sonarservice.SonarService{
		Name:      "ldap",
		Port:      389,
		Transport: sonarservice.TCP,
	}
	testCases["ldap_389"] = sonarservice.SonarService{
		Name:      "ldap",
		Port:      389,
		Transport: sonarservice.NotDef,
	}
	testCases["udp_ldap_389"] = sonarservice.SonarService{
		Name:      "ldap",
		Port:      389,
		Transport: sonarservice.UDP,
	}
	testCases["udp_ldap_starttls_389"] = sonarservice.SonarService{
		Name:      "ldap_starttls",
		Port:      389,
		Transport: sonarservice.UDP,
	}
	testCases["gtp-c_2123"] = sonarservice.SonarService{
		Name:      "gtp-c",
		Port:      2123,
		Transport: sonarservice.NotDef,
	}

	for k, v := range testCases {
		out, err := sonarservice.NewSonarService(k)
		rq.NoError(err)
		rq.Equal(out, v)
	}
}
