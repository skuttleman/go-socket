package services

import (
  "fmt"
  "github.com/skuttleman/tilda-poc/Godeps/_workspace/src/github.com/googollee/go-socket.io"
)

func Socket() *socketio.Server {
  server, _ := socketio.NewServer(nil)
  server.On("connection", func(so socketio.Socket) {
    so.Join("chat")
    fmt.Println("connect:", so)
    so.On("chat message", func(msg string) {
      // fmt.Println("emit:", so.Emit("chat message", msg))
      so.BroadcastTo("chat", "chat message", msg)
    })
    so.On("disconnection", func() {
      fmt.Println("disconnect:", so)
    })
  })
  server.On("error", func(so socketio.Socket, err error) {
    fmt.Println("error:", err)
  })
  return server
}
