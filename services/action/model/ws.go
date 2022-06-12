package model

import (
	"github.com/gorilla/websocket"
	"time"
)

const (
	WriteWait      = 10 * time.Second
	PongWait       = 60 * time.Second
	PingPeriod     = (PongWait * 9) / 10
	MaxMessageSize = 512
)

type Connection struct {
	WsConn   *websocket.Conn
	Send     chan []byte
	LimitNum int
	TimeLog  int64
}

type Message struct {
	Data   []byte
	RoomId string
	Name   string
	Conn   *Connection
}

type Hub struct {
	Rooms        map[string]map[*Connection]bool
	Broadcast    chan Message
	BroadcastALl chan Message
	Register     chan Message
	Unregister   chan Message
	Kickoutroom  chan Message
}

var Manager = Hub{
	Broadcast:    make(chan Message),
	BroadcastALl: make(chan Message),
	Register:     make(chan Message),
	Unregister:   make(chan Message),
	Kickoutroom:  make(chan Message),
	Rooms:        make(map[string]map[*Connection]bool),
}

type User struct {
	Id       int
	Name     string
	Password string
}
