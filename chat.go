package chatnio

import "github.com/Deeptrain-Community/chatnio-api-go/utils"

type Chat struct {
	Id    int
	Uri   string
	Token string
	Conn  *utils.WebSocket
}

type ChatAuthForm struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

type ChatRequestForm struct {
	Message string `json:"message"`
	Model   string `json:"model"`
	Web     bool   `json:"web"`
}

type ChatPartialResponse struct {
	Message string  `json:"message"`
	Keyword string  `json:"keyword"`
	Quota   float32 `json:"quota"`
	End     bool    `json:"end"`
}

func (i *Instance) NewChat(id int) *Chat {
	return &Chat{
		Id:    id,
		Uri:   i.GetChatEndpoint(),
		Token: i.GetApiKey(),
		Conn:  utils.NewWebsocket(i.GetChatEndpoint()),
	}
}

func (c *Chat) Send(v interface{}) bool {
	return c.Conn.Send(v)
}

func (c *Chat) Close() error {
	return c.Conn.Close()
}

func (c *Chat) SendAuthRequest() bool {
	return c.Send(ChatAuthForm{
		Id:    c.Id,
		Token: c.Token,
	})
}

func (c *Chat) AskStream(form *ChatRequestForm, callback func(ChatPartialResponse)) {
	// for authentication
	if c.Conn.IsEmpty() {
		c.SendAuthRequest()
	}

	c.Send(form)
	for {
		form := utils.ReadForm[ChatPartialResponse](c.Conn)
		if form == nil {
			continue
		}

		callback(*form)
		if form.End {
			break
		}
	}
}

func (c *Chat) Ask(form *ChatRequestForm, channel chan ChatPartialResponse) {
	c.AskStream(form, func(res ChatPartialResponse) {
		channel <- res
	})
}
