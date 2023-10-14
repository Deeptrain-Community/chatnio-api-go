package chatnio

import (
	"fmt"
	"github.com/Deeptrain-Community/chatnio-api-go/utils"
	"os"
	"strings"
)

type Instance struct {
	ApiKey   string
	Endpoint string
}

func (i *Instance) GetApiKey() string {
	return i.ApiKey
}

func (i *Instance) GetEndpoint() string {
	if i.Endpoint == "" {
		return "https://api.chatnio.net"
	}
	return i.Endpoint
}

func (i *Instance) SetApiKey(apiKey string) {
	i.ApiKey = apiKey
}

func (i *Instance) SetEndpoint(endpoint string) {
	if strings.HasSuffix(endpoint, "/") {
		endpoint = strings.TrimSuffix(endpoint, "/")
	}
	i.Endpoint = endpoint
}

func (i *Instance) GetChatApiKey() string {
	if !i.IsAuthenticated() {
		return "anonymous"
	}
	return i.ApiKey
}

func (i *Instance) IsAuthenticated() bool {
	return strings.TrimSpace(i.ApiKey) != ""
}

func (i *Instance) GetChatEndpoint() string {
	host := utils.TrimPrefixes(i.GetEndpoint(), "http://", "https://")
	return fmt.Sprintf("wss://%s/chat", host)
}

func (i *Instance) GetHeaders() utils.Headers {
	return utils.Headers{
		"Authorization": fmt.Sprintf("Bearer %s", i.GetApiKey()),
		"Content-Type":  "application/json",
		"Accept":        "application/json",
	}
}

func (i *Instance) Mix(path string) string {
	return i.GetEndpoint() + path
}

// NewInstance creates a new instance of the chatnio client
func NewInstance(key string) *Instance {
	return &Instance{
		ApiKey: key,
	}
}

// NewInstanceFromEnv creates a new instance of the chatnio client from the environment
func NewInstanceFromEnv(env string) *Instance {
	key := os.Getenv(env)
	return NewInstance(key)
}
