package main

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {

	http.HandleFunc("/websocket", handleConnections)

	fmt.Println("http server started on :8000")
	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		fmt.Printf("ListenAndServe: %v\n", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		glog.Fatal(err)
	}

	defer ws.Close()

	closeChan := make(chan int)

	defer close(closeChan)

	go func() {
		for {
			_, msg, err := ws.ReadMessage()
			if err != nil {
				closeChan <- 1
				return
			}
			fmt.Printf("recv message: %s\n", string(msg))
		}
	}()

	tc := time.Tick(2 * time.Second)
	for i := 1; i <= 200; i++ {

		select {
		case <-tc:
			fmt.Printf("send count: %d\n", i)

			msg := fmt.Sprintf("send count: %d", i)
			err := ws.WriteMessage(websocket.TextMessage, []byte(msg))

			if err != nil {
				fmt.Printf("Failed to send message to client!")
				return
			}
		case <-closeChan:
			fmt.Println("Conn is closed by client.")
			return
		}

	}
}
