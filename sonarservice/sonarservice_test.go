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

func TestSonarServiceToString(t *testing.T) {
	rq := require.New(t)
	testCases := make(map[sonarservice.SonarService]string)

	s1 := sonarservice.SonarService{
		Name:      "ldap",
		Port:      389,
		Transport: sonarservice.TCP,
	}
	testCases[s1] = "tcp_ldap_389"

	s2 := sonarservice.SonarService{
		Name:      "ldap",
		Port:      389,
		Transport: sonarservice.NotDef,
	}
	testCases[s2] = "ldap_389"

	s3 := sonarservice.SonarService{
		Name:      "ldap",
		Port:      389,
		Transport: sonarservice.UDP,
	}
	testCases[s3] = "udp_ldap_389"

	s4 := sonarservice.SonarService{
		Name:      "ldap_starttls",
		Port:      389,
		Transport: sonarservice.UDP,
	}
	testCases[s4] = "udp_ldap_starttls_389"

	s5 := sonarservice.SonarService{
		Name:      "gtp-c",
		Port:      2123,
		Transport: sonarservice.NotDef,
	}
	testCases[s5] = "gtp-c_2123"

	for service, serviceName := range testCases {
		name := service.ToString()
		rq.Equal(name, serviceName)
	}
}
