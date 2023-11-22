package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type User struct {
	Name    string `json:"Name"`
	Message string `json:"Message"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients []websocket.Conn

func main() {
	http.HandleFunc("/echo", wsHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.ListenAndServe(":8080", nil)
}
