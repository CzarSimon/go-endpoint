package endpoint

import (
	"fmt"
)

// ServerAddr default values
const (
	DEFAULT_SERVER_PROTOCOL = "http"
	DEFAULT_SERVER_HOST     = "localhost"
)

// ServerAddr Connection info for a server
type ServerAddr struct {
	Host     string `json:"host,omitempty"`
	Port     string `json:"port,omitempty"`
	Protocol string `json:"protocol,omitempty"`
}

// NewServerAddr Creates a new ServerAddr by checking
// environment varialbes whose keys are derived from the supplied name
func NewServerAddr(name string) ServerAddr {
	return ServerAddr{
		Host:     getenv(makeKey(name, HOST), DEFAULT_SERVER_HOST),
		Port:     getenv(makeKey(name, PORT), ""),
		Protocol: getenv(makeKey(name, PROTOCOL), DEFAULT_SERVER_PROTOCOL),
	}
}

// ToURL Accepts a route as a parameter and turns the route + ServerAddr to a full url
func (server ServerAddr) ToURL(route string) string {
	return fmt.Sprintf(
		"%s://%s%s/%s", server.Protocol, server.Host, server.getPortString(), route)
}

// Network Gets the ServerAddr protwork based on the protocol
func (server ServerAddr) Network() string {
	return mapProtocolToNetwork(server.Protocol)
}

// Address Gets the host:port concatentaion of the ServerAddr
func (server ServerAddr) Address() string {
	return fmt.Sprintf("%s%s", server.Host, server.getPortString())
}

func (server ServerAddr) getPortString() string {
	if server.Port == "" {
		return ""
	}
	return fmt.Sprintf(":%s", server.Port)
}

func mapProtocolToNetwork(protocol string) string {
	switch protocol {
	case "udp":
		return "udp"
	case "unix":
		return "unix"
	default:
		return "tcp"
	}
}
