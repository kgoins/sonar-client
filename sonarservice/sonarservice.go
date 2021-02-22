package sonarservice

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type SonarService struct {
	Name      string
	Port      int
	Transport TransportProtocol
}

// extractPort takes a split service string and returns
// the port and remaining string entries
func extractPort(service []string) ([]string, int, error) {
	port, err := strconv.Atoi(service[len(service)-1])
	if err != nil {
		return nil, port, err
	}

	return service[:len(service)-1], port, nil
}

// extractTransport will find and remove the transport
// definition from the service array, returning the remainder
func extractTransport(service []string) ([]string, TransportProtocol) {
	tp, err := NewTransportProtocol(service[0])
	if err != nil {
		tp = NotDef
		// no transport present, return as is
		return service, tp
	}

	return service[1:], tp
}

func NewSonarService(serviceStr string) (SonarService, error) {
	var emptyService SonarService

	serviceArr := strings.Split(serviceStr, "_")
	if len(serviceArr) < 2 {
		return emptyService, errors.New("Invalid service string")
	}

	serviceArr, port, err := extractPort(serviceArr)
	if err != nil {
		return emptyService, err
	}

	serviceArr, tp := extractTransport(serviceArr)
	serviceName := strings.Join(serviceArr, "_")

	return SonarService{
		Name:      serviceName,
		Port:      port,
		Transport: tp,
	}, nil
}

// ToString produces a Sonar file compliant string
// representation of the service
func (s SonarService) ToString() string {
	portStr := strconv.Itoa(s.Port)
	if s.Transport == NotDef {
		return fmt.Sprintf(
			"%s_%s",
			s.Name,
			portStr,
		)
	}

	return fmt.Sprintf(
		"%s_%s_%s",
		string(s.Transport),
		s.Name,
		portStr,
	)
}
