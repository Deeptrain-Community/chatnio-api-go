package utils

import (
	"github.com/gorilla/websocket"
	"io"
)

type WebSocket struct {
	Conn   *websocket.Conn
	Cursor int
}

func NewWebsocket(url string) (ws *WebSocket, err error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return
	}

	ws = &WebSocket{
		Conn:   conn,
		Cursor: 0,
	}
	return
}

func (w *WebSocket) GetCursor() int {
	return w.Cursor
}

func (w *WebSocket) IncrCursor() {
	w.Cursor++
}

func (w *WebSocket) IsEmpty() bool {
	return w.Cursor == 0
}

func (w *WebSocket) Read() (int, []byte, error) {
	return w.Conn.ReadMessage()
}

func (w *WebSocket) Write(messageType int, data []byte) error {
	w.IncrCursor()
	return w.Conn.WriteMessage(messageType, data)
}

func (w *WebSocket) Close() error {
	return w.Conn.Close()
}

func (w *WebSocket) DeferClose() {
	if err := w.Close(); err != nil {
		return
	}
}

func (w *WebSocket) NextWriter(messageType int) (io.WriteCloser, error) {
	return w.Conn.NextWriter(messageType)
}

func (w *WebSocket) ReadJSON(v interface{}) error {
	return w.Conn.ReadJSON(v)
}

func (w *WebSocket) SendJSON(v interface{}) error {
	w.IncrCursor()
	return w.Conn.WriteJSON(v)
}

func (w *WebSocket) Send(v interface{}) bool {
	return w.SendJSON(v) == nil
}

func (w *WebSocket) Receive(v interface{}) bool {
	return w.ReadJSON(v) == nil
}

func (w *WebSocket) SendText(message string) bool {
	return w.Write(websocket.TextMessage, []byte(message)) == nil
}

func ReadForm[T interface{}](w *WebSocket) *T {
	// golang cannot use generic type in class-like struct
	// except ping
	_, message, err := w.Read()
	if err != nil {
		return nil
	} else if string(message) == "{\"type\":\"ping\"}" {
		return ReadForm[T](w)
	}

	form, err := UnmarshalByteForm[T](message)
	if err != nil {
		return nil
	}
	return form
}
