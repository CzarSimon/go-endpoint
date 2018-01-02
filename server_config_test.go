package endpoint

import "testing"

func TestToURL(t *testing.T) {
	serverWithPort := ServerConfig{
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
	serverWithoutPort := ServerConfig{
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

func TestNewServerConfig(t *testing.T) {
	setEnvForTest(TEST_NAME)
	server := NewServerConfig(TEST_NAME)
	if server.Host != EXPECTED_HOST {
		t.Errorf("NewServerConfig failed. Expected Host=%s Got=%s",
			EXPECTED_HOST, server.Host)
	}
	if server.Port != EXPECTED_PORT {
		t.Errorf("NewServerConfig failed. Expected Port=%s Got=%s",
			EXPECTED_PORT, server.Port)
	}
	if server.Protocol != EXPECTED_PROTOCOL {
		t.Errorf("NewServerConfig failed. Expected Protocol=%s Got=%s",
			EXPECTED_PROTOCOL, server.Protocol)
	}
	server = NewServerConfig("NOT_PRESENT_NAME")
	if server.Host != DEFAULT_SERVER_HOST {
		t.Errorf("NewServerConfig failed. Expected Host=%s Got=%s",
			EXPECTED_HOST, server.Host)
	}
	if server.Port != "" {
		t.Errorf("NewServerConfig failed. Expected Port=\"\" Got=%s", server.Port)
	}
	if server.Protocol != DEFAULT_SERVER_PROTOCOL {
		t.Errorf("NewServerConfig failed. Expected Protocol=%s Got=%s",
			DEFAULT_SERVER_PROTOCOL, server.Protocol)
	}
}
