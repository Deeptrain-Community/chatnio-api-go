package utils

import (
	"testing"
)

func TestNewWebsocket(t *testing.T) {
	ws := NewWebsocket("wss://api.chatnio.net/chat")
	if ws == nil {
		t.Error("websocket is nil")
	}
}
