package utils

import (
	"fmt"
	"testing"
)

func TestNewWebsocket(t *testing.T) {
	_, err := NewWebsocket("wss://api.chatnio.net/chat")
	if err != nil {
		t.Errorf(fmt.Sprintf("error occurred: %s", err.Error()))
	}
}
