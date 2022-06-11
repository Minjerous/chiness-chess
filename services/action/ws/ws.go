package ws

import (
	"github.com/gorilla/websocket"
	"time"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

type connection struct {
	ws            *websocket.Conn
	send          chan []byte
	limitNum      int
	forbiddenWord bool
	timeLog       int64
}

type message struct {
	Data   []byte
	RoomId string
	Name   string
	Conn   *connection
}

type hub struct {
	Rooms        map[string]map[*connection]bool //房间
	broadCast    chan message
	broadCastALL chan message
	register     chan message
	unregister   chan message
	kickoutroom  chan message
	warnmsg      chan message
}

var h = hub{
	broadCast:    make(chan message),
	broadCastALL: make(chan message),
	warnmsg:      make(chan message),
	register:     make(chan message),
	unregister:   make(chan message),
	kickoutroom:  make(chan message),
	//rooms:        make(map[string]map[*connection]bool),
}

type User struct {
	Id       int
	Name     string
	Password string
}
