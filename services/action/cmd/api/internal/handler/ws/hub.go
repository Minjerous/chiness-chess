package ws

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	NewGrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		//跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)
