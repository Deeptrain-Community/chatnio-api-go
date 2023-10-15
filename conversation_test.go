package chatnio

import (
	"fmt"
	"testing"
)

func TestInstance_GetConversationList(t *testing.T) {
	if !instance.IsAuthenticated() {
		return
	}

	list, err := instance.GetConversationList()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(fmt.Sprintf("[conversation] list: %v", list))
}

func TestInstance_GetConversation(t *testing.T) {
	if !instance.IsAuthenticated() {
		return
	}

	conv, err := instance.GetConversation(2)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(fmt.Sprintf("[conversation] load: %v", conv))
}

func TestInstance_DeleteConversation(t *testing.T) {
	if true {
		return
	} // // delete conversation is disabled because it's unsafe (it will delete conversation and then the test will fail)
	if !instance.IsAuthenticated() {
		return
	}

	err := instance.DeleteConversation(2)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(fmt.Sprintf("[conversation] delete: %v", err))
}
