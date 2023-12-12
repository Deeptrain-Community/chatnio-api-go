package chatnio

import (
	"fmt"
	"github.com/Deeptrain-Community/chatnio-api-go/utils"
	"testing"
)

func TestChat_AskStream(t *testing.T) {
	if !instance.IsAuthenticated() {
		return
	}

	chat, err := instance.NewChat(-1)
	if err != nil {
		t.Errorf(fmt.Sprintf("error occurred: %s", err.Error()))
		return
	}

	defer chat.DeferClose()

	chat.AskStream(&ChatRequestForm{
		Message: "hello",
		Model:   "gpt-3.5-turbo",
	}, func(resp ChatPartialResponse) {
		fmt.Println(fmt.Sprintf("[chat] ask stream: %s", utils.MarshalForm(resp)))
	})
}

func TestChat_Ask(t *testing.T) {
	if !instance.IsAuthenticated() {
		return
	}

	chat, err := instance.NewChat(-1)
	defer chat.DeferClose()

	if err != nil {
		t.Errorf(fmt.Sprintf("error occurred: %s", err.Error()))
		return
	}

	channel := make(chan ChatPartialResponse)
	chat.Ask(&ChatRequestForm{
		Message: "hello",
		Model:   "gpt-3.5-turbo",
	}, channel)

	for resp := range channel {
		fmt.Println(fmt.Sprintf("[chat] ask: %s", utils.MarshalForm(resp)))
	}
}
