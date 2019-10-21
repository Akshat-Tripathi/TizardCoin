package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	accounts = loadAccounts()
	word     = "dhtlzhccsd"
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients = hub{
		make(map[string]*websocket.Conn),
	}
)

func main() {
	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("css/"))))
	http.Handle("/static/js/", http.StripPrefix("/static/js/", http.FileServer(http.Dir("js/"))))
	http.Handle("/static/imgs/", http.StripPrefix("/static/imgs/", http.FileServer(http.Dir("imgs/"))))
	http.HandleFunc("/login/", loginHandler)
	http.HandleFunc("/", pageHandler)
	http.ListenAndServe("localhost:8080", nil)
}
