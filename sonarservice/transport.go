package sonarservice

import "errors"

type TransportProtocol string

const (
	TCP    TransportProtocol = "tcp"
	UDP    TransportProtocol = "udp"
	NotDef TransportProtocol = ""
)

func NewTransportProtocol(s string) (TransportProtocol, error) {
	t := TransportProtocol(s)
	switch t {
	case TCP, UDP, NotDef:
		return t, nil
	}

	return t, errors.New("Invalid transport protocol")
}
