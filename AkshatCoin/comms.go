package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

//Info - communications happen using infos
type Info struct {
	Word string
	Num  int
}

type hub struct {
	users map[string]*websocket.Conn
}

func (h *hub) sendToAll(data Info) {
	for _, v := range clients.users {
		fmt.Println(v.WriteJSON(data))
	}
}

func (h *hub) sendToClient(data Info, client string) {
	user := clients.users[client]
	fmt.Println(user.WriteJSON(data))
}
