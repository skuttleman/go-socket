package services

import (
  "fmt"
  "github.com/skuttleman/tilda-poc/Godeps/_workspace/src/github.com/googollee/go-socket.io"
)

type Game struct {
  Name string
  Player1 socketio.Socket
  Player2 socketio.Socket
}

type Connections struct {
  Waiting socketio.Socket
  Games []Game
}

var gameId uint64

func Socket() *socketio.Server {
  server, _ := socketio.NewServer(nil)

  server.On("connection", func(so socketio.Socket) {
    name := "Room" + fmt.Sprint(gameId / 2)
    gameId++
    so.Join(name)
    so.BroadcastTo(name, "game move", "connected")
    so.On("game move", func(msg string) {
      so.BroadcastTo(name, "game move", msg)
    })
    so.On("disconnection", func() {
      server.BroadcastTo(name, "disconnect")
    })
  })
  // server.On("error", func(so socketio.Socket, err error) {
  //   fmt.Println("error:", err)
  // })
  return server
}
