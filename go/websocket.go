package main

import (
	"fmt"
	"golang.org/x/net/websocket"
        "log"
	//"encoding/base64"
	"encoding/base64"
)

var url = "wss://kubernetes:6443/api/v1/namespaces/default/pods/aur2-h6zf6/exec?stdout=1&stdin=1&stderr=1&tty=1&command=%2Fbin%2Fsh&command=-i"
var origin = "https://kubernetes:6443"
func main() {
	ws, err := websocket.Dial(url, "base64.channel.k8s.io", origin)
	if err != nil {
		log.Fatal(err)
	}
	message := []byte("0")
	_, err = ws.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", message)

	var first = true
	var msg = make([]byte, 512)
	for {
		m, err := ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}

		s, _ := base64.StdEncoding.DecodeString(string(msg[1:m]))
		fmt.Printf("Receive: %s\n", s)

		if (string(msg[:m]) == "1IyA=" && first) {
			first = false
			l := base64.StdEncoding.EncodeToString([]byte("ls\n"))
			fmt.Printf("Send: %s\n", l)
			ws.Write([]byte("0" + l))
		}
	}

	ws.Close()
}
