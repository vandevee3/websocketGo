package main

import (
	"fmt"
	"net/http"
	"sync"
)

func doPrint(wg *sync.WaitGroup, message string, i interface{}) {
	defer wg.Done()
	fmt.Printf("%s sent: %s\n", i, string(message))
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	var wg sync.WaitGroup
	clients = append(clients, *conn)

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		wg.Add(1)
		go doPrint(&wg, string(msg), conn.RemoteAddr())

		for _, client := range clients {
			if err = client.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
		wg.Wait()
	}

}
