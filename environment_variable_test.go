package endpoint

import (
	"os"
	"testing"
)

const (
	TEST_NAME         = "ENDPOINT_UNIT_TESTING"
	EXPECTED_HOST     = "endpoint-unit-test"
	EXPECTED_PORT     = "1337"
	EXPECTED_PROTOCOL = "https"
)

func TestMakeKey(t *testing.T) {
	key := makeKey("BACKEND", PORT)
	exprectedKey := "BACKEND_PORT"
	if key != exprectedKey {
		t.Errorf("makeKey() failed. Expected=%s Got=%s", exprectedKey, key)
	}
	key = makeKey("BACKEND", HOST)
	exprectedKey = "BACKEND_HOST"
	if key != exprectedKey {
		t.Errorf("makeKey() failed. Expected=%s Got=%s", exprectedKey, key)
	}
}

func TestGetenv(t *testing.T) {
	setEnvForTest(TEST_NAME)
	result := getenv(makeKey(TEST_NAME, HOST), DEFAULT_SERVER_HOST)
	if result != EXPECTED_HOST {
		t.Errorf("getenv() main branch failed. Expected=%s Got=%s",
			EXPECTED_HOST, result)
	}
	result = getenv(makeKey("NOT_PRESENT_NAME", HOST), DEFAULT_SERVER_HOST)
	if result != DEFAULT_SERVER_HOST {
		t.Errorf("getenv() default value failed. Expected=%s Got=%s",
			DEFAULT_SERVER_HOST, result)
	}
}

func setEnvForTest(name string) {
	os.Setenv(makeKey(name, HOST), EXPECTED_HOST)
	os.Setenv(makeKey(name, PORT), EXPECTED_PORT)
	os.Setenv(makeKey(name, PROTOCOL), EXPECTED_PROTOCOL)
}
