package endpoint

import (
	"fmt"
)

const (
	DEFAULT_SERVER_PROTOCOL = "http"
	DEFAULT_SERVER_HOST     = "localhost"
)

// ServerConfig Connection info for a server
type ServerConfig struct {
	Host     string `json:"host,omitempty"`
	Port     string `json:"port,omitempty"`
	Protocol string `json:"protocol,omitempty"`
}

// NewServerConfig Creates a new ServerConfig by checking
// environment varialbes whose keys are derived from the supplied name
func NewServerConfig(name string) ServerConfig {
	return ServerConfig{
		Host:     getenv(makeKey(name, HOST), DEFAULT_SERVER_HOST),
		Port:     getenv(makeKey(name, PORT), ""),
		Protocol: getenv(makeKey(name, PROTOCOL), DEFAULT_SERVER_PROTOCOL),
	}
}

// ToURL Accepts a route as a parameter and turns the route + ServerConfig to a full url
func (server ServerConfig) ToURL(route string) string {
	port := server.Port
	if port != "" {
		port = fmt.Sprintf(":%s", server.Port)
	}
	return fmt.Sprintf("%s://%s%s/%s", server.Protocol, server.Host, port, route)
}
