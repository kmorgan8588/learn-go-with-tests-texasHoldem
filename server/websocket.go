package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type playerServerWS struct {
	*websocket.Conn
}

func NewPlayerServerWS(w http.ResponseWriter, r *http.Request) *playerServerWS {
	conn, err := wsUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("problem upgrading connectoin to Websockets %v\n", err)
	}

	return &playerServerWS{conn}
}

func (w *playerServerWS) WaitForMsg() string {
	_, msg, err := w.ReadMessage()
	if err != nil {
		log.Printf("error reading message from websocket %v\n", err)
	}
	return string(msg)
}