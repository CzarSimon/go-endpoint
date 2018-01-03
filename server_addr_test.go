package endpoint

import "testing"

func TestToURL(t *testing.T) {
	serverWithPort := ServerAddr{
		Host:     "localhost",
		Port:     "3000",
		Protocol: "http",
	}
	expectedURL := "http://localhost:3000/test"
	resultURL := serverWithPort.ToURL("test")
	if expectedURL != resultURL {
		t.Errorf("ServerConfig.ToURL() failed. Expected=%s Got=%s",
			expectedURL, resultURL)
	}
	serverWithoutPort := ServerAddr{
		Host:     "localhost",
		Protocol: "http",
	}
	expectedURL = "http://localhost/test"
	resultURL = serverWithoutPort.ToURL("test")
	if expectedURL != resultURL {
		t.Errorf("ServerConfig.ToURL() failed. Expected=%s Got=%s",
			expectedURL, resultURL)
	}
}

func TestNewServerAddr(t *testing.T) {
	setEnvForTest(TEST_NAME)
	server := NewServerAddr(TEST_NAME)
	if server.Host != EXPECTED_HOST {
		t.Errorf("NewServerAddr failed. Expected Host=%s Got=%s",
			EXPECTED_HOST, server.Host)
	}
	if server.Port != EXPECTED_PORT {
		t.Errorf("NewServerAddr failed. Expected Port=%s Got=%s",
			EXPECTED_PORT, server.Port)
	}
	if server.Protocol != EXPECTED_PROTOCOL {
		t.Errorf("NewServerAddr failed. Expected Protocol=%s Got=%s",
			EXPECTED_PROTOCOL, server.Protocol)
	}
	server = NewServerAddr("NOT_PRESENT_NAME")
	if server.Host != DEFAULT_SERVER_HOST {
		t.Errorf("NewServerAddr failed. Expected Host=%s Got=%s",
			EXPECTED_HOST, server.Host)
	}
	if server.Port != "" {
		t.Errorf("NewServerAddr failed. Expected Port=\"\" Got=%s", server.Port)
	}
	if server.Protocol != DEFAULT_SERVER_PROTOCOL {
		t.Errorf("NewServerAddr failed. Expected Protocol=%s Got=%s",
			DEFAULT_SERVER_PROTOCOL, server.Protocol)
	}
}

func TestServerAddrAddress(t *testing.T) {
	serverWithPort := getServerAddrWithPort()
	exprectedAddress := "localhost:3000"
	if exprectedAddress != serverWithPort.Address() {
		t.Errorf("ServerAddr.Address() failed. Expected=%s Got=%s",
			exprectedAddress, serverWithPort.Address())
	}
}

func TestServerAddrNetwork(t *testing.T) {
	serverWithPort := getServerAddrWithPort()
	exprectedNetwork := "tcp"
	if exprectedNetwork != serverWithPort.Network() {
		t.Errorf("ServerAddr.Network() failed. Expected=%s Got=%s",
			exprectedNetwork, serverWithPort.Network())
	}
	serverWithPort.Protocol = "unix"
	exprectedNetwork = "unix"
	if exprectedNetwork != serverWithPort.Network() {
		t.Errorf("ServerAddr.Network() failed. Expected=%s Got=%s",
			exprectedNetwork, serverWithPort.Network())
	}
	serverWithPort.Protocol = "udp"
	exprectedNetwork = "udp"
	if exprectedNetwork != serverWithPort.Network() {
		t.Errorf("ServerAddr.Network() failed. Expected=%s Got=%s",
			exprectedNetwork, serverWithPort.Network())
	}
}

func getServerAddrWithPort() ServerAddr {
	return ServerAddr{
		Host:     "localhost",
		Port:     "3000",
		Protocol: "http",
	}
}
