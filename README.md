# ChatNio Go Library

The official Go library for the Chat Nio API

- Authors: Deeptrain Team
- Free software: MIT license 
- Documentation: https://docs.chatnio.net

## Features

- Chat 
- Conversation 
- Quota 
- Subscription and Package

## Installation

```shell
go get -u github.com/Deeptrain-Community/chatnio-api-go
```

## Usage

- Authentication
```go
instance := chatnio.NewInstance("sk-...")
// or load from env
instance := chatnio.NewInstanceFromEnv("CHATNIO_TOKEN")
```

- Chat
```go
chat := instance.NewChat(-1) // id -1: creating new chat

// using hook
chat.AskStream(&chatnio.ChatRequestForm{
	Message: "hello",
	Model:   "gpt-3.5-turbo",
}, func(resp chatnio.ChatPartialResponse) {
	fmt.Println(resp.End, resp.Quota, resp.Message, resp.Keyword)
})

// or using channel
channel := make(chan chatnio.ChatPartialResponse)
chat.Ask(&chatnio.ChatRequestForm{
    Message: "hello",
    Model:   "gpt-3.5-turbo",
}, channel)
for resp := range channel {
    // do something
}
```

- Conversation
```go
// list conversations
if list, err := instance.GetConversationList(); err == nil {
    for _, conv := range list {
        fmt.Println(conv.Id, conv.Name)
    }
}

// get conversation
if conv, err := instance.GetConversation(1); err == nil {
    fmt.Println(conv.Id, conv.Name, conv.Message)
}

// delete conversation
if err := instance.DeleteConversation(1); err == nil {
    fmt.Println("deleted")
}
```

- Quota
```go
// get quota
if quota, err := instance.GetQuota(); err == nil {
    fmt.Println(fmt.Sprintf("current quota is %0.2f", quota))
}

// buy quota
if err := instance.BuyQuota(100); err == nil {
    fmt.Println("bought successfully")
}
```

- Subscription and Package
```go
// get subscription
if sub, err := instance.GetSubscription(); err == nil {
    fmt.Println(sub.IsSubscribed, sub.Expired)
}

// buy subscription
if err := instance.Subscribe(1, 1); err == nil {
    fmt.Println("bought basic subscription for 1 month successfully")
}

// get package
if pkg, err := instance.GetPackage(); err == nil {
    fmt.Println(pkg.Data.Cert, pkg.Data.Teenager)
}
```


## Test
To run the tests, you need to set the environment variable CHATNIO_TOKEN to your secret key.

```shell
export CHATNIO_TOKEN=sk-...
```

Then run the tests:

```shell
go test -v ./...
```
