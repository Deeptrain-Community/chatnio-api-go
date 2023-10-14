package chatnio

import (
	"testing"
)

var instance *Instance

func init() {
	instance = NewInstanceFromEnv("CHATNIO_TOKEN")
}

func TestInstance_GetEndpoint(t *testing.T) {
	if instance.GetEndpoint() != "https://api.chatnio.net" {
		t.Error("endpoint is not https://api.chatnio.net")
	}
}

func TestInstance_GetChatEndpoint(t *testing.T) {
	if instance.GetChatEndpoint() != "wss://api.chatnio.net/chat" {
		t.Error("chat endpoint is not correct")
	}
}

func TestInstance_GetHeaders(t *testing.T) {
	headers := instance.GetHeaders()
	if headers["Authorization"] != "Bearer "+instance.GetApiKey() {
		t.Error("authorization header is not correct")
	}
}

func TestInstance_Mix(t *testing.T) {
	if instance.Mix("/test") != "https://api.chatnio.net/test" {
		t.Error("mix is not correct")
	}
}

func TestInstance_IsAuthenticated(t *testing.T) {
	if len(instance.ApiKey) > 0 && !instance.IsAuthenticated() {
		t.Error("authentication is not correct")
	}
}

func TestInstance_GetChatApiKey(t *testing.T) {
	if instance.IsAuthenticated() && instance.GetChatApiKey() != instance.ApiKey {
		t.Error("chat api key is not correct")
	} else if !instance.IsAuthenticated() && instance.GetChatApiKey() != "anonymous" {
		t.Error("chat api key is not correct")
	}
}
