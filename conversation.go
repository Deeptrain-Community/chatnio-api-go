package chatnio

import (
	"fmt"
	"github.com/Deeptrain-Community/chatnio-api-go/utils"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Conversation struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Message []Message `json:"message"`
}

type ConversationList struct {
	Status  bool           `json:"status"`
	Message string         `json:"message"`
	Data    []Conversation `json:"data"`
}

type ConversationLoad struct {
	Status  bool         `json:"status"`
	Message string       `json:"message"`
	Data    Conversation `json:"data"`
}

type ConversationDelete struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func (i *Instance) GetConversationList() ([]Conversation, error) {
	data, err := utils.GetForm[ConversationList](i.Mix("/conversation/list"), i.GetHeaders())

	if err != nil {
		return nil, err
	} else if !data.Status {
		return nil, fmt.Errorf(data.Message)
	}

	return data.Data, nil
}

func (i *Instance) GetConversation(id int) (*Conversation, error) {
	data, err := utils.GetForm[ConversationLoad](i.Mix(fmt.Sprintf("/conversation/load?id=%d", id)), i.GetHeaders())

	if err != nil {
		return nil, err
	} else if !data.Status {
		return nil, fmt.Errorf(data.Message)
	}

	return &data.Data, nil
}

func (i *Instance) DeleteConversation(id int) error {
	data, err := utils.GetForm[ConversationDelete](i.Mix(fmt.Sprintf("/conversation/delete?id=%d", id)), i.GetHeaders())

	if err != nil {
		return err
	} else if !data.Status {
		return fmt.Errorf(data.Message)
	}

	return nil
}
