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
go get -u github.com/deeptrain-community/chatnio-api-go
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
