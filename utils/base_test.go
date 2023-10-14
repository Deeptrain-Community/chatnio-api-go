package utils

import "testing"

func TestTrimPrefixes(t *testing.T) {
	data := "https://api.chatnio.net"
	expected := "api.chatnio.net"
	actual := TrimPrefixes(data, "http://", "https://")
	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}
